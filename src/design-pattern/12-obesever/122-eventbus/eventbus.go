package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (bus *AsyncEventBus) Subscribe(topic string, function interface{}) error {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	value := reflect.ValueOf(function)
	if value.Type().Kind() != reflect.Func {
		return fmt.Errorf("handlers is not a function")
	}

	handlers, ok := bus.handlers[topic]
	if !ok {
		handlers = []reflect.Value{}
	}
	bus.handlers[topic] = append(handlers, value)

	return nil
}

// Publish 发布
// 这里异步执行，并且不会等待返回结果
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]

	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	for i, _ := range handlers {
		go handlers[i].Call(params)
	}
}
