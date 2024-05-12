package parser

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type SectionDTO struct {
	Title string
	Body  string
}

func createDoc(url string) goquery.Document {

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

	return *doc
}

func parseUrl(doc goquery.Document, titleSelect string, bodySelect string) SectionDTO {
	section := SectionDTO{}

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
