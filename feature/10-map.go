package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 3

	fmt.Println("m:", m)
	fmt.Println("v1:", m["k1"])

	delete(m, "k1")
	fmt.Println("m:", m)

	_, k1_is_present := m["k1"]
	_, k2_is_present := m["k2"]
	fmt.Println("k1_is_present:", k1_is_present)
	fmt.Println("k2_is_present:", k2_is_present)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("n:", n)
}
