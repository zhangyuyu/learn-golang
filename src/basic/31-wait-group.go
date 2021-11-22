package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker2(&wg)
	}

	wg.Wait()
	fmt.Println(total2)
}

func worker1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

var total struct {
	sync.Mutex
	value int
}

var total2 uint64

func worker2(wg *sync.WaitGroup) {
	defer wg.Done()
	var i uint64
	for i = 0; i <= 100; i++ {
		atomic.AddUint64(&total2, i)
	}
}

// WaitGroup 用于等待该函数启动的所有协程。
