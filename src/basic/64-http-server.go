package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8090", nil)
}

func health(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	fmt.Println("server: health Handler started")
	defer fmt.Println("server: health Handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(writer, "OK\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
