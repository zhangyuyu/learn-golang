package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommandFunc(t *testing.T) {
	// 模拟客户端事件：一个线程不断接受客户端命令
	eventChan := make(chan string)

	go func() {
		events := []string{"start", "archive", "start", "archive"}
		for _, event := range events {
			eventChan <- event
		}
	}()
	defer close(eventChan)

	// 使用命令队列缓存命令：另外一个线程丢进队列，进行处理
	commands := make(chan Command, 1000)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command Command
			switch event {
			case "start":
				command = StartCommandFunc()
			case "archive":
				command = ArchiveCommandFunc()
			}

			// 将命令入队
			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
