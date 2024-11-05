package main

import (
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you are %s\n", c.Query("name"), c.Path)
	})

	engine.GET("/yyj", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<name>yyj<name>")
	})

	engine.GET("/hcp", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"name": "yyj",
			"lover": "hcp",
		})
	})

	engine.Run(":9999")
}
