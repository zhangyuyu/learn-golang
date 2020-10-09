package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.ParseFloat("1.234", 64))
	fmt.Println(strconv.ParseInt("123", 0, 64))
	fmt.Println(strconv.ParseInt("0x1c8", 0, 64))
	fmt.Println(strconv.ParseUint("789", 0, 64))

	fmt.Println(strconv.Atoi("135"))
	fmt.Println(strconv.Atoi("wat"))
}
