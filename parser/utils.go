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

func saveImageFromUrlToImages(fullUrl string) string {
	response, e := http.Get(fullUrl)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	filename := buildFileNameFromUrl(fullUrl)
	//open a file for writing
	file, err := os.Create("./images/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
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
