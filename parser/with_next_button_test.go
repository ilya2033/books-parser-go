package parser

import (
	"os"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestAdd(t *testing.T) {
	buttonSelect := "a.btn.next_page"
	reader, err := os.Open("./test/nextButtonTest.html")

	if err != nil {
		t.Errorf("File parse error")
	}

	doc, err := goquery.NewDocumentFromReader(reader)

	if err != nil {
		t.Errorf("Reader parse error")
	}

	got, err := parseNextButtonUrl(*doc, buttonSelect)
	want := "http://ncorrect.com"

	if err != nil {
		t.Errorf("Error while reading button")
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
