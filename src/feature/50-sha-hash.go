package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// New Write Sum
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil) // 参数用来追加额外的字节切片

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
