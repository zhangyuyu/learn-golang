package command

import (
	"fmt"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
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
	commands := make(chan ICommand, 1000)
	defer close(commands)

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}

			var command ICommand
			switch event {
			case "start":
				command = &StartCommand{}
			case "archive":
				command = &ArchiveCommand{}
			}

			// 将命令入队
			commands <- command
		}
	}()

	for {
		select {
		case c := <-commands:
			c.Execute()
		case <-time.After(1 * time.Second):
			fmt.Println("timeout 1s")
			return
		}
	}
}
