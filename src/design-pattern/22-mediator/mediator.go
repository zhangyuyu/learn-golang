package mediator

import (
	"fmt"
)

// Mediator 中介者接口
type Mediator interface {
	operation(receiver int, msg string)
	register(receiver int, colleague Colleague)
}

type ConcreteMediator struct {
	colleagues map[int]Colleague
}

func (c *ConcreteMediator) operation(receiver int, msg string) {
	colleague, ok := c.colleagues[receiver]
	if ok {
		colleague.sendMsg(msg)
	} else {
		fmt.Println("no matched receiver!")
	}
}

func (c *ConcreteMediator) register(receiver int, colleague Colleague) {
	if c.colleagues == nil {
		c.colleagues = make(map[int]Colleague)
	}
	c.colleagues[receiver] = colleague
}

type Colleague interface {
	startDialog(msg string)
	sendMsg(msg string)
}

type ConcreteColleagueA struct {
	mediator Mediator
}

func (c *ConcreteColleagueA) startDialog(msg string) {
	c.mediator.operation(2, msg)
}

func (c *ConcreteColleagueA) sendMsg(msg string) {
	fmt.Println("ConcreteColleagueA send msg: " + msg)
}

type ConcreteColleagueB struct {
	mediator Mediator
}

func (c *ConcreteColleagueB) startDialog(msg string) {
	c.mediator.operation(1, msg)
}

func (c *ConcreteColleagueB) sendMsg(msg string) {
	fmt.Println("ConcreteColleagueB send msg: " + msg)
}
