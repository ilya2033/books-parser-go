package parser

import (
	"fmt"
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

func TestRemoveScripts(*testing.T) {
	testCases := []struct {
		expected string
		value    string
	}{
		{value: "some random text <sicript></script>", expected: "some random text "},
		{value: "<script>someando text </script>", expected: ""},
		{value: "<tr><td></td><script>sdad</script></tr>", expected: "<tr><td></td></tr>"},
	}

	for _, tc := range testCases {

		got := RemoveScripts(tc.value)

		if got != tc.expected {
			fmt.Errorf("got %q, wanted %q", got, tc.expected)
		}

	}
}
