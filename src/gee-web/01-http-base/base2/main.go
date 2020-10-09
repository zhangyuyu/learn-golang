package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":8090", engine))
}

type Engine struct{}

func (engine Engine) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/":
		fmt.Fprintf(writer, "URL.path = %q\n", request.URL.Path)
	case "/hello":
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(writer, "404 NOT FOUND: %s\n", request.URL)
	}
}
