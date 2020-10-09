package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go workder(done)
	<-done
}

func workder(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
