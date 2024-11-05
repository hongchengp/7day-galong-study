package main

import "net/http"

type Engine struct {}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/hello":
		w.Write([]byte("hello http"))
		break
	case "/yyj":
		w.Write([]byte("yyj ai hcp"))
		break
	default:
		w.Write([]byte("url error"))
	}
}

func main() {
	http.ListenAndServe(":9999", &Engine{})
}