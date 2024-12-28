package parser

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const PATH_TO_IMAGES = "./images/"

func saveImageFromUrlToImages(fullUrl string) string {
	response := getImageFromUrl(fullUrl)
	defer response.Body.Close()

	filename := buildFileNameFromUrl(fullUrl)

	resultFileName := saveFromReader(filename, response.Body, PATH_TO_IMAGES)
	return resultFileName
}

func getImageFromUrl(url string) *http.Response {
	response, e := http.Get(url)

	if e != nil {
		log.Fatal(e)
	}

	return response
}

func saveFromReader(filename string, reader io.ReadCloser, path string) string {
	file, err := os.Create(path + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		log.Fatal(err)
	}

	return file.Name()
}

func buildFileNameFromUrl(fullUrl string) string {
	fileUrl, err := url.Parse(fullUrl)
	if err != nil {
		log.Fatalln(err)
	}

	path := fileUrl.Path
	segments := strings.Split(path, "/")

	fileName := segments[len(segments)-1]

	return (fileName)
}

func RemoveScripts(html string) string {
	scriptRe := regexp.MustCompile(`(?s)<script.*<\/script>`)
	endScriptRe := regexp.MustCompile(`</script>`)

	result := scriptRe.ReplaceAllStringFunc(html, func(s string) string {
		index := endScriptRe.FindStringIndex(s)
		return strings.TrimSpace(s[index[1]:])
	})
	return result
}
