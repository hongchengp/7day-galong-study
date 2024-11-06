package gee

import (
	"strings"
)

type router struct {
	handlers map[string]HandlerFunc
	roots map[string]*node
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
		roots: make(map[string]*node),
	}
}

func parsePattern(pattern string) []string {
	Arr := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range Arr {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (router *router) addRoute(method string, pattern string, handler HandlerFunc) {
	root := router.roots[method]
	if root == nil {
		root = newNode("")
		router.roots[method] = root
	}
	root.insert(pattern, parsePattern(pattern), 0)

	key := method + "-" + pattern
	router.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string){
	searchParts := parsePattern(path)
	root := r.roots[method]
	if root == nil {
		return nil, nil
	}
	n := root.search(searchParts, 0)
	if n == nil {
		return nil, nil
	}

	parts := parsePattern(n.pattern)
	params := make(map[string]string)
	for index, part := range parts {
		if part[0] == ':' {
			params[part[1:]] = searchParts[index]
		} else if part[0] == '*' {
			params[part[1:]] = strings.Join(searchParts[index:], "/")
		}
	}

	return n, params
}

// func (router *router) handle(c *Context) {
// 	n, params := router.getRoute(c.Method, c.Path)
// 	if n == nil {
// 		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
// 		return
// 	}

// 	c.Params = params 
// 	key := c.Method + "-" + n.pattern
// 	router.handlers[key](c)
// }