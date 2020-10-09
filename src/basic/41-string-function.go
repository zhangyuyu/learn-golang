package main

import (
	"fmt"
	"strings"
)

func main() {
	p := fmt.Println
	str := "This is an example of a string"

	// 前缀
	prefix := "Th"
	fmt.Printf("Does the string '%s' have prefix '%s'?\n", str, prefix)
	p(strings.HasPrefix(str, prefix))

	// 后缀
	suffix := "string"
	fmt.Printf("Does the string '%s' have suffix '%s'?\n", str, suffix)
	p(strings.HasSuffix(str, suffix))

	// 包含
	substr := "example"
	fmt.Printf("Does the string '%s' contains '%s'?\n", str, substr)
	p(strings.Contains(str, substr))

	// 位置
	fmt.Printf("The position of '%s' is:\n", substr)
	p(strings.Index(str, substr))

	p("Count:    ", strings.Count("test", "t"))
	p("Join:     ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:   ", strings.Repeat("he", 5))
	p("Replace:  ", strings.Replace("foo", "o", "+", -1))
	p("Replace:  ", strings.Replace("foo", "o", "+", 1))
	p("Split:    ", strings.Split("a, b, c, d", ","))
	p("Split:    ", strings.Split("a,b,c,d", ","))
	p("ToLower:  ", strings.ToLower("TEST"))
	p("ToUpper:  ", strings.ToUpper("test"))
	p("Len:      ", len("hello"))
	p("Char:     ", "hello"[0])
	p("Char:     ", "a"[0])
	p("Char:     ", 'a')
	p("Char:     ", "a")
}
