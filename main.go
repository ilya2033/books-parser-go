package main

import (
	"ilya2033/book-parser/parser"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/parse-from-url-list", fromUrls)
	router.POST("/parse-with-next-button", withNextButton)
	router.Run(":3001")
}

func fromUrls(c *gin.Context) {
	var json parser.MultiUrlParserSettings

	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(400, err.Error())
	}

	epub := parser.StartParsingMultiUrl(json)
	// handle error
	err = epub.Write("./epub/" + json.Title + ".epub")

	if err != nil {
		c.JSON(500, err.Error())
	}

	removeGlob("./images")
}

func withNextButton(c *gin.Context) {
	var json parser.NextButtonParserSettings

	err := c.BindJSON(&json)
	if err != nil {
		c.JSON(400, err.Error())
	}

	epub := parser.StartParsingWithNextButton(json)
	// handle error
	err = epub.Write("./epub/" + json.Title + ".epub")

	if err != nil {
		c.JSON(500, err.Error())
	}

	removeGlob("./images/*")
}

func removeGlob(path string) (err error) {
	contents, err := filepath.Glob(path)
	if err != nil {
		return
	}

	for _, item := range contents {
		err = os.RemoveAll(item)
		if err != nil {
			return
		}
	}
	return
}
