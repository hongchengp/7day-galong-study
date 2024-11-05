package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	W http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context{
	return &Context{
		W: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.W.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.W.WriteHeader(code)
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.W.Header().Set("Content-Type", "text/plain")
	c.Status(code)	
	c.W.Write([]byte(fmt.Sprintf(format, values...)))
}

func (c *Context) JSON(code int,obj interface{}) {
	c.W.Header().Set("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.W)
	encoder.Encode(obj)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.W.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.W.Header().Set("Content-Type", "text/html")
	c.Status(code) 
	c.W.Write([]byte(html))
}

