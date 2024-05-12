package parser

import (
	"github.com/bmaupin/go-epub"
)

type MultiUrlParserSettings struct {
	Urls        []string `json:"urls" binding:"required"`
	TitleSelect string   `json:"title-select" binding:"required"`
	BodySelect  string   `json:"body-select" binding:"required"`
	Author      string   `json:"author" binding:"required"`
	Title       string   `json:"title" binding:"required"`
}

func StartParsingMultiUrl(settings MultiUrlParserSettings) *epub.Epub {
	book := epub.NewEpub(settings.Title)
	book.SetAuthor(settings.Author)

	for _, value := range settings.Urls {
		doc := createDoc(value)
		section := parseUrl(doc, settings.TitleSelect, settings.BodySelect)
		body := "<h1>" + section.Title + "</h1>" + "<p>" + section.Body + "</p>"
		book.AddSection(body, section.Title, "", "")
	}

	return book
}
