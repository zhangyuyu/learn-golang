package gee

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (r *router) addRouter(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(ctx *Context) {
	key := ctx.Method + "-" + ctx.Path
	handler := r.handlers[key]
	if handler != nil {
		handler(ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}
