package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	hello()

	countCharcters("asSASA ddd dsjkdsjs dk")
	countCharcters("asSASA ddd dsjkdsjs dk こん")
}

func countCharcters(str string) {
	// count number of characters:
	fmt.Printf("The number of bytes in '%s' is %d\n", str, len(str))
	fmt.Printf("The number of characters in '%s' is %d\n", str, utf8.RuneCountInString(str))
}

func hello() {
	println("Hello", "world")
}
