package main

import (
	"fmt"
)

func main() {
	a, b := getValues(12)
	fmt.Println(a, b)

	_, c := getValues(15)
	fmt.Println(c)
}

func getValues(num int) (int, int) {
	return num / 10, num % 10
}
