package parser

import (
	"testing"
)

func TestBuildFileNameFromUrl(t *testing.T) {
	testCases := []struct {
		expected string
		value    string
	}{
		{value: "https://test.com/some/path/to/book_name_123", expected: "book_name_123"},
		{value: "https://test.com/book_name_123", expected: "book_name_123"},
		{value: "https://test.com/book-name-123", expected: "book-name-123"},
	}

	for _, tc := range testCases {

		got := buildFileNameFromUrl(tc.value)

		if got != tc.expected {
			t.Errorf("got %q, wanted %q", got, tc.expected)
		}

	}
}
