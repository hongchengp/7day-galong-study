package gee

import (
	"net/http"
	"strings"
)

type HandlerFunc func(c *Context)


type RouterGroup struct {
	prefix string
	middlewares []HandlerFunc
	engine *Engine
}

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	routerGroup := &RouterGroup{
		engine: engine,
	}
	engine.RouterGroup = routerGroup
	engine.groups = []*RouterGroup{routerGroup}
	return engine
}

func (r *RouterGroup) Group(prefix string) *RouterGroup{
	engine := r.engine
	routerGroup := &RouterGroup{
		engine: engine,
	}
	fullPrefix := r.prefix + prefix 
	routerGroup.prefix = fullPrefix
	engine.groups = append(engine.groups, routerGroup)
	return routerGroup
}

func (r *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	fullPattern := r.prefix + pattern
	engine := r.engine
	engine.router.addRoute(method, fullPattern, handler)
}

func (r *RouterGroup) GET(pattern string, handler HandlerFunc) {
	r.addRoute("GET", pattern, handler)
}

func (r *RouterGroup) POST(pattern string, handler HandlerFunc) {
	r.addRoute("POST", pattern, handler)
}

func (r *RouterGroup) Use(Handlers ...HandlerFunc) {
	if r.middlewares == nil {
		r.middlewares = make([]HandlerFunc, 0)
	}
	r.middlewares = append(r.middlewares, Handlers...)
}

func (engine *Engine) Run(addr string) {
	http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	n, params := engine.router.getRoute(c.Method, c.Path)
	if n == nil {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		return
	}

	c.Params = params 
	pattern := n.pattern
	for _, group := range engine.groups {
		if strings.HasPrefix(pattern, group.prefix) {
			c.handlers = append(c.handlers, group.middlewares...)
		}
	}
	key := req.Method + "-" + pattern
	c.handlers = append(c.handlers, engine.router.handlers[key])
	c.Next()
}