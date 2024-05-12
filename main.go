package main

import (
	"fmt"
	"ilya2033/book-parser/parser"

	"github.com/gin-gonic/gin"
)

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

	fmt.Println(epub.Title())
}

func main() {
	router := gin.Default()
	router.POST("/parse-from-url-list", fromUrls)
	router.Run(":3001")
}
