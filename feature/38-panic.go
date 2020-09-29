package main

import (
	"os"
)

func main() {
	panic("a problem")

	_, err := os.Open("feature/01-hello.go")

	if err != nil {
		panic(err)
	}
}
