package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dict:", b)

	var c [2][3]int
	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[0]); j++ {
			c[i][j] = i + j
		}
	}
	fmt.Println("2d:", c)
}
