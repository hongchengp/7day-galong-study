package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "yyj aaa hao\n")
	})

	g1 := r.Group("v1")
	g1.GET("/hello/:name", func(c *gee.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s, yyj ai hcp\n", name)
	})

	g1.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello http\n")
	})

	g2 := g1.Group("/yyj")
	g2.GET("/lala", func(c *gee.Context) {
		c.String(http.StatusOK, "啦啦啦啦\n")
	})
	r.Run(":9999")
}
