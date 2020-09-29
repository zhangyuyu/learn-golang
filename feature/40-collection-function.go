package main

import (
	"fmt"
	"strings"
)

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(Map(strs, strings.ToUpper))
}

func Map(strs []string, f func(string) string) []string {
	strm := make([]string, len(strs))
	for i, v := range strs {
		strm[i] = f(v)
	}
	return strm
}

func Filter(strs []string, f func(string) bool) []string {
	strm := make([]string, 0)

	for _, v := range strs {
		if f(v) {
			strm = append(strm, v)
		}
	}
	return strm
}

func Any(strs []string, f func(v string) bool) bool {
	for _, v := range strs {
		if f(v) {
			return true
		}
	}
	return false
}

func All(strs []string, f func(v string) bool) bool {
	for _, v := range strs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Include(strs []string, s2 string) bool {
	return Index(strs, s2) >= 0
}

func Index(strs []string, s2 string) int {
	for i, v := range strs {
		if v == s2 {
			return i
		}
	}
	return -1
}
