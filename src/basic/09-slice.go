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

// 测试扩容的原理
func testSliceExpand() {
	arr := []int{1}

	myfunc1(arr)
	fmt.Println(arr) // [1]

	arr = append(arr, 3)
	arr = append(arr, 4)
	myfunc2(arr)
	fmt.Println(arr) // [9 3 4]，不是 [1 3 4]，因为原容量还够，扩容之后指向的是原数组
}

func myfunc1(arr []int) {
	arr = append(arr, 2)
	arr[0] = 0
	fmt.Println(arr) // [0, 2]
	return
}

func myfunc2(arr []int) {
	arr = append(arr, 5)
	arr[0] = 9
	fmt.Println(arr) // [9 3 4 5]
	return
}
