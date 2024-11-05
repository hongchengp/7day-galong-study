package main

import "net/http"

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello http"))
} 

func love(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("yyj ai hcp"))
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/yyj", love)

	http.ListenAndServe(":9999", nil)
}