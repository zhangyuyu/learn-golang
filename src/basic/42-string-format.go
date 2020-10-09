package main

import (
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	p := point{1, 2}

	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%p\n", &p) // 指针的值

	fmt.Printf("%T\n", p)
	fmt.Printf("%t\n", true) // 格式化布尔值

	fmt.Printf("%d\n", 123)
	fmt.Printf("%b\n", 8)  // 输出二进制表示形式
	fmt.Printf("%c\n", 33) // 输出给定整数的对应字符

	fmt.Printf("%x\n", 456)        // 十六进制
	fmt.Printf("%x\n", "hex this") // 十六进制

	fmt.Printf("%f\n", 78.9)        // 浮点型
	fmt.Printf("%e\n", 123400000.0) // 科学计数法
	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")
	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("|%6d|%6d|\n", 12, 345) // 指定宽度，右对齐
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45) // 指定宽度，左对齐
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("hello %s", "world")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
