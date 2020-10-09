package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"GolangHelloworld/src/utils"
)

func main() {
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	fmt.Println(filepath.IsAbs("dir/filename"))
	fmt.Println(filepath.IsAbs("/dir/filename"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel，找相对路径
	rel, err := filepath.Rel("a/b", "a/b/t/filename")
	utils.Check(err)
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/filename")
	utils.Check(err)
	fmt.Println(rel)
}
