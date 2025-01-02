package parser

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/bmaupin/go-epub"
)

type SelectParserSettings struct {
	Url              string `json:"url" binding:"required"`
	TitleSelect      string `json:"title-select" binding:"required"`
	BodySelect       string `json:"body-select" binding:"required"`
	SelectListSelect string `json:"select-list-select" binding:"required"`
	Author           string `json:"author" binding:"required"`
	Title            string `json:"title" binding:"required"`
	ConverUrl        string `json:"cover-url" binding:"required"`
}

const MILLISECONDS_WAIT__SELECT = 300

func StartParsingWithSelect(settings SelectParserSettings) *epub.Epub {
	var err error
	book := epub.NewEpub(settings.Title)
	book.SetAuthor(settings.Author)

	url := settings.Url
	counter := 0

	doc := createDocFromUrl(url)
	urls, err := parseAllSelectUrls(doc, settings.SelectListSelect)

	if urls == nil || err != nil {
		log.Fatalln(err)
	}

	for _, url := range urls {
		doc = createDocFromUrl(url)
		addSectionToEpub(book, doc, settings)

		time.Sleep(MILLISECONDS_WAIT__SELECT * time.Millisecond)

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

func addSectionToEpub(book *epub.Epub, doc goquery.Document, settings SelectParserSettings) {
	section := parseSection(doc, settings.TitleSelect, settings.BodySelect)
	body := "<h1>" + section.Title + "</h1>" + "<p>" + section.Body + "</p>"
	book.AddSection(body, section.Title, "", "")
}

func parseAllSelectUrls(doc goquery.Document, selectListSelect string) ([]string, error) {
	selectUrls := make([]string, 0)

	doc.Find(selectListSelect).Each(func(i int, s *goquery.Selection) {
		url, ok := s.Attr("href")

		if ok {
			selectUrls = append(selectUrls, url)
		}
	})

	if len(selectUrls) == 0 {
		return nil, errors.New("Can not find select urls")
	}

	return selectUrls, nil
}
