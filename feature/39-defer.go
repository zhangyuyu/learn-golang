package main

import (
	"fmt"
	"os"
)

func main() {
	f := creatFile("defer.txt")

	defer closeFile(f)
	writeFile(f)
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "some data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()
	if err != nil {
		fmt.Println(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func creatFile(p string) *os.File {
	fmt.Println("creating")

	f, err := os.Create(p)

	if err != nil {
		panic(err)
	}
	return f
}
