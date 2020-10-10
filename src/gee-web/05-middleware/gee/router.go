package gee

import (
	"log"
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %s - %s", method, pattern)
	values := parsePattern(pattern)
	key := method + "-" + pattern
	root := r.roots[method]
	if root == nil {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, values, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchValues := parsePattern(path)
	root := r.roots[method]
	if root == nil {
		return nil, nil
	}

	params := make(map[string]string)
	child := root.search(searchValues, 0)

	if child != nil {
		values := parsePattern(child.pattern)
		for index, value := range values {
			if value[0] == ':' {
				params[value[1:]] = searchValues[index]
			}
			if value[0] == '*' && len(value) > 1 {
				params[value[1:]] = strings.Join(searchValues[index:], "/")
				break
			}
		}
		return child, params
	}

	return nil, nil
}

func (r *router) handle(ctx *Context) {
	root, params := r.getRoute(ctx.Method, ctx.Path)

	if root != nil {
		ctx.Params = params
		key := ctx.Method + "-" + root.pattern
		handler := r.handlers[key]
		ctx.handlers = append(ctx.handlers, handler)
	} else {
		ctx.handlers = append(ctx.handlers, func(context *Context) {
			ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
		})
	}

	ctx.Next()
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	values := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			values = append(values, item)
			// Only one * is allowed
			if item[0] == '*' {
				break
			}
		}
	}
	return values
}
