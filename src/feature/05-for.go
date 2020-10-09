package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for k := 0; k <= 5; k++ {
		if k%2 == 0 {
			continue
		}
		fmt.Println(k)
	}
}
