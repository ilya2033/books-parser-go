package parser

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseNextButtonUrl(t *testing.T) {
	buttonSelect := "a.btn.next_page"
	reader, err := os.Open("./test/page_1.html")

	if err != nil {
		t.Errorf("File parse error")
	}

	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		t.Errorf("Reader parse error")
	}

	got, err := parseNextButtonUrl(*doc, buttonSelect)
	want := "http://correct.com"

	if err != nil {
		t.Errorf("Error while reading button")
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
