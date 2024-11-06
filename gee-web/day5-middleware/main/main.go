package main

import (
	"fmt"
	"gee"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "yyj aaa hao\n")
	})

	g1 := r.Group("v1")
	g1.Use(func(c *gee.Context) {
		// Start timer
		t := time.Now()
		log.Printf("[%d] %s in %v for group v1", c.StatusCode, c.Req.RequestURI, time.Since(t))
		c.Next()
	})
	g1.Use(func(c *gee.Context) {
		fmt.Println("yyj ai hcp")
		c.Next()
	})
	g1.GET("/hello/:name", func(c *gee.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s, yyj ai hcp\n", name)
		fmt.Println("before...")
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
