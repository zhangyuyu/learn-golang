package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "This is an example of a string"

	// 前缀
	prefix := "Th"
	fmt.Printf("Does the string '%s' have prefix '%s'?\n", str, prefix)
	fmt.Println(strings.HasPrefix(str, prefix))

	// 后缀
	suffix := "string"
	fmt.Printf("Does the string '%s' have suffix '%s'?\n", str, suffix)
	fmt.Println(strings.HasSuffix(str, suffix))

	// 包含
	substr := "example"
	fmt.Printf("Does the string '%s' contains '%s'?\n", str, substr)
	fmt.Println(strings.Contains(str, substr))

	// 位置
	fmt.Printf("The position of '%s' is:\n", substr)
	fmt.Println(strings.Index(str, substr))

}
