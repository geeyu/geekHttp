package gee

import (
	"net/http"
)

type router struct {
	Handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		Handlers: make(map[string]HandlerFunc),
	}
}

func (router *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.Handlers[key] = handler
}

func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := router.Handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
