package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "b", "c"}
	fmt.Println("s:", s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("s:", s)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
