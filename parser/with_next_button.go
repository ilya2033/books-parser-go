package parser

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bmaupin/go-epub"
)

type NextButtonParserSettings struct {
	Url              string `json:"url" binding:"required"`
	TitleSelect      string `json:"title-select" binding:"required"`
	BodySelect       string `json:"body-select" binding:"required"`
	NextButtonSelect string `json:"next-button-select" binding:"required"`
	Author           string `json:"author" binding:"required"`
	Title            string `json:"title" binding:"required"`
	ConverUrl        string `json:"cover-url" binding:"required"`
}

const MILLISECONDS_WAIT = 300

func StartParsingWithNextButton(settings NextButtonParserSettings) *epub.Epub {
	var err error
	book := epub.NewEpub(settings.Title)
	book.SetAuthor(settings.Author)

	url := settings.Url
	counter := 0

	for {
		doc := createDoc(url)

		section := parseSection(doc, settings.TitleSelect, settings.BodySelect)
		body := "<h1>" + section.Title + "</h1>" + "<p>" + section.Body + "</p>"
		book.AddSection(body, section.Title, "", "")

		url, err = parseNextButtonUrl(doc, settings.NextButtonSelect)

		if url == "" || err != nil {
			break
		}
		time.Sleep(MILLISECONDS_WAIT * time.Millisecond)

		log.Println(fmt.Sprintf("Page: %d", counter))
		counter++
	}

	coverImage, err := book.AddImage(saveImageFromUrlToImages(settings.ConverUrl), "coverImage")
	if err != nil {
		log.Fatalln(err)
	}

	book.SetCover(coverImage, "")

	return book
}

func parseNextButtonUrl(doc goquery.Document, nextButtonSelect string) (string, error) {
	buttonUrls := make([]string, 0)

	doc.Find(nextButtonSelect).Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")

		if ok {
			buttonUrls = append(buttonUrls, url)
		}
	})

	if len(buttonUrls) == 0 {
		return "", errors.New("Can not find next button url")
	}

	return buttonUrls[0], nil
}
