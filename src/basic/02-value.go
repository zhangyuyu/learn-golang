package main

import (
	_ "embed"
	"fmt"
)

//go:embed 00-README.md
var some string

func main() {
	fmt.Println(some)
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(!true)
}
