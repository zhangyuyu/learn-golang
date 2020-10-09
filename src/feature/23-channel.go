package main

import (
	"fmt"
)

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "pang"
		messages <- "ending"
		close(messages) // 必须要，否则会死锁
	}()

	msg, ok := <-messages
	fmt.Println(msg, ok)

	for i := range messages {
		fmt.Println(i)
	}

	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("Done")
}
