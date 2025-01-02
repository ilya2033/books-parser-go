package parser

import (
	"bytes"
	"ilya2033/book-parser/parser/test/mocks"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestParseSection(t *testing.T) {
	testCases := []struct {
		expected      SectionDTO
		value         string
		bodySelector  string
		titleSelector string
	}{
		{
			// Case 1: Empty HTML string and empty selectors
			value: "",
			expected: SectionDTO{
				Title: "",
				Body:  "",
			},
			titleSelector: "",
			bodySelector:  "",
		},
		{
			// Case 2: Valid HTML with selectors matching content
			value: `<html><head><title>Test Title</title></head><body><div id="content"><p>Test Body</p></div></body></html>`,
			expected: SectionDTO{
				Title: "Test Title",
				Body:  "<p>Test Body</p>",
			},
			titleSelector: "title",
			bodySelector:  "#content",
		},
		{
			// Case 3: Valid HTML but selectors do not match
			value: `<html><head><title>Test Title</title></head><body><div id="content">Test Body</div></body></html>`,
			expected: SectionDTO{
				Title: "",
				Body:  "",
			},
			titleSelector: ".nonexistent",
			bodySelector:  ".nonexistent",
		},
		{
			// Case 4: HTML with missing title or body content
			value: `<html><head></head><body><div id="content"><p>Test Body</p></div></body></html>`,
			expected: SectionDTO{
				Title: "",
				Body:  "<p>Test Body</p>",
			},
			titleSelector: "title",
			bodySelector:  "#content",
		},
		{
			// Case 5: HTML with nested elements, retaining only formatting tags
			value: `<html><head><title>Nested Title</title></head><body><div id="content"><p>Nested <strong>Body</strong></p></div></body></html>`,
			expected: SectionDTO{
				Title: "Nested Title",
				Body:  "<p>Nested <strong>Body</strong></p>",
			},
			titleSelector: "title",
			bodySelector:  "#content",
		},
	}

	for _, tc := range testCases {
		reader := &mocks.ReadCloserMock{
			Data: bytes.NewReader([]byte(tc.value)),
		}
		doc, err := goquery.NewDocumentFromReader(reader)

		if err != nil {
			t.Error(err)
		}

		got := parseSection(*doc, tc.titleSelector, tc.bodySelector)

		if got != tc.expected {
			t.Errorf("got %q, wanted %q", got, tc.expected)
		}

	}
}
