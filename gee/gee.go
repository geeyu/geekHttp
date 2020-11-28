package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

func (engine *Engine) Get(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

func (engine *Engine) Post(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := newContext(w, r)
	engine.router.handle(ctx)
}
