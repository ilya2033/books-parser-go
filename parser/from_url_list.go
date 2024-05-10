package parser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/bmaupin/go-epub"
)

type MultiUrlParserSettings struct {
	Urls        []string `json:"urls" binding:"required"`
	TitleSelect string   `json:"title-select" binding:"required"`
	BodySelect  string   `json:"body-select" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Title       string   `json:"title" binding:"required"`
}

type SectionDTO struct {
	Title string
	Body  string
}

func StartParsing(settings MultiUrlParserSettings) *epub.Epub {
	book := epub.NewEpub(settings.Title)
	book.SetAuthor(settings.Author)

	for _, value := range settings.Urls {
		section := parseUrl(value, settings.TitleSelect, settings.BodySelect)
		body := "<h1>" + section.Title + "</h1>" + "<p>" + section.Body + "</p>"
		book.AddSection(body, section.Title, "", "")
	}

	return book
}

func parseUrl(url string, titleSelect string, bodySelect string) SectionDTO {
	section := SectionDTO{}

	res, err := http.Get(url)

	if err != nil {
		log.Fatalln(err.Error())
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	titleHtml, err := doc.Find(titleSelect).First().Html()
	if err != nil {
		log.Fatalln(err.Error())
	}

	bodyHtml, err := doc.Find(bodySelect).First().Html()
	if err != nil {
		log.Fatalln(err.Error())
	}
	section.Title = RemoveScripts(titleHtml)
	section.Body = RemoveScripts(bodyHtml)

	return section
}
