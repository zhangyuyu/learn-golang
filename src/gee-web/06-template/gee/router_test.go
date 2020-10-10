package gee

import (
	"reflect"
	"testing"
)

func Test_parsePattern(t *testing.T) {
	type args struct {
		pattern string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"参数", args{"/p/:name"}, []string{"p", ":name"}},
		{"一个通配符", args{"/p/*"}, []string{"p", "*"}},
		{"多个通配符", args{"/p/*name/*"}, []string{"p", "*name"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func Test_router_getRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/geektutu")

	type result struct {
		pattern string
		arg     string
	}
	want := &result{"/hello/:name", "geektutu"}

	if !reflect.DeepEqual(n.pattern, want.pattern) {
		t.Errorf("getRoute() pattern = %v, want %v", n.pattern, want.pattern)
	}
	if !reflect.DeepEqual(ps["name"], want.arg) {
		t.Errorf("getRoute() name = %v, want %v", ps["name"], want.arg)
	}
}
