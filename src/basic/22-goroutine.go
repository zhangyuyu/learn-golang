package main

import (
	"fmt"
	"time"
)

func main() {
	counter("direct")

	go counter("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("anonymous...")

	time.Sleep(time.Second)
	fmt.Println("done")
}

func counter(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}
