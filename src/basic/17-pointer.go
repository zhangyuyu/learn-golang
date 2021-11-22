package main

import (
	"fmt"
)

func main() {
	//testPinter()
	testPointer()
}

func testPinter() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func zeroval(ival int) {
	ival = 0
}

func testPointer() {
	nums := []int{1, 2}
	twiceSlice(nums)
	fmt.Print(nums)

	sliceHeader := IntSliceHeader{
		Data: []int{1, 2},
		Len:  2,
		Cap:  2,
	}
	twice(sliceHeader)
	fmt.Print(sliceHeader.Data)
}

func twiceSlice(x []int) {
	for i := range x {
		x[i] *= 2
	}
}

type IntSliceHeader struct {
	Data []int
	Len  int
	Cap  int
}

func twice(x IntSliceHeader) {
	for i := 0; i < x.Len; i++ {
		x.Data[i] *= 2
	}
}
