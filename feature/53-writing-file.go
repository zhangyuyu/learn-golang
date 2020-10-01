package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("writing-dat1.txt", d1, 0644)
	check(err)

	f, err := os.Create("writing-dat2.txt")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync() // 调用 Sync 将缓冲区的数据写入硬盘。

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer。
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
