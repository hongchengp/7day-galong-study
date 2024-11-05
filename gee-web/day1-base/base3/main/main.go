package main

import(
	"gee"
	"net/http"
)

func main() {
	engine := gee.New()
	engine.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello http"))
	})
	engine.GET("/yyj", func (w http.ResponseWriter, r *http.Request)  {
		w.Write([]byte("yyj ai hcp"))
	})

	engine.Run(":9999")
}
