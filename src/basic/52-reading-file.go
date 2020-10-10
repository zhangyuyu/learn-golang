package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"learn-golang/src/utils"
)

func main() {
	dat, err := ioutil.ReadFile("defer.txt")
	utils.Check(err)
	fmt.Println(string(dat))

	f, err := os.Open("defer.txt")
	utils.Check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	utils.Check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek 到一个文件中已知的位置，并从这个位置开始读取。
	o2, err := f.Seek(6, 0)
	utils.Check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	utils.Check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	o3, err := f.Seek(6, 0)
	utils.Check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	utils.Check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	utils.Check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	utils.Check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // 通常这个操作应该在OPen操作之后立即使用defer来完成
}
