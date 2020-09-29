package main

import (
	"fmt"
)

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// 把pings的消息，放在pongs里
func pong(pings <-chan string, pongs chan string) {
	msg := <-pings
	pongs <- msg
}

// 把要发送的消息，放在pings
func ping(pings chan<- string, msg string) {
	pings <- msg
}
