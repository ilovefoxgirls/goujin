package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handleIndex(c *gin.Context) {
	c.HTML(200, "page/index.tmpl", gin.H{
		"title": "Index page",
	})
}

func main() {
	fmt.Println("test")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*/*.tmpl")

	router.GET("/", handleIndex)
	router.GET("/:id/", func(c *gin.Context) {
		rid := c.Param("id")
		rpage := c.Param("page")
		msg := "id: " + rid + " page: " + rpage
		c.String(200, msg)
	})
	router.GET("/:id/:page/", func(c *gin.Context) {
		rid := c.Param("id")
		rpage := c.Param("page")
		msg := "id: " + rid + " page: " + rpage
		c.String(200, msg)
	})
	router.Run(":8080")
}
