package eventbus

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func handler1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("function1, %s+%s\n", msg1, msg2)
}

func handler2(msg1, msg2 string) {
	fmt.Printf("function2, %s-%s\n", msg1, msg2)
}

func TestAsyncEventBus_Publish(t *testing.T) {
	bus := NewAsyncEventBus()
	err := bus.Subscribe("topic:1", handler1)
	assert.NoError(t, err)
	err = bus.Subscribe("topic:1", handler2)
	assert.NoError(t, err)
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}
