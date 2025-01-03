package parser

import (
	"bytes"
	"ilya2033/book-parser/parser/test/mocks"
	"os"
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

func TestRemoveScripts(t *testing.T) {
	testCases := []struct {
		expected string
		value    string
	}{
		{value: "some random text <script></script>", expected: "some random text "},
		{value: "<script>someando text </script>", expected: ""},
		{value: "<tr><td></td><script>sdad</script></tr>", expected: "<tr><td></td></tr>"},
	}

	for _, tc := range testCases {

		got := RemoveScripts(tc.value)

		if got != tc.expected {
			t.Errorf("got %q, wanted %q", got, tc.expected)
		}

	}
}

func TestSaveFromReader(t *testing.T) {
	reader := &mocks.ReadCloserMock{
		Data: bytes.NewReader([]byte("BlaBlaBla")),
	}
	filename := "test___blablalba.png"
	expectedFilePath := "./test/" + filename
	os.Remove(expectedFilePath)
	defer os.Remove(expectedFilePath)

	got := saveFromReader(filename, reader, "./test/")

	if got != expectedFilePath {
		t.Errorf("got %q, wanted %q", got, filename)
	}

	_, err := os.Stat(expectedFilePath)

	if err != nil {
		t.Errorf("file dose not exists")
	}

}
