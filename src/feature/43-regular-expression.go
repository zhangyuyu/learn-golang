package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	pattern := "p([a-z]+ch)"
	isMatch, _ := regexp.MatchString(pattern, "peach")
	fmt.Println(isMatch)

	r, _ := regexp.Compile(pattern)
	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))         // 查找首次匹配的字符串，匹配开始和结束位置的索引
	fmt.Println(r.FindStringSubmatch("peach punch"))      // 完全匹配和局部匹配的字符串
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // 完全匹配和局部匹配的字符串

	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllString("peach punch pinch", 2))
	fmt.Println(r.FindAllStringIndex("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatch("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile(pattern)
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))
	out := r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)
	fmt.Println(string(out))
}
