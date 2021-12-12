package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {
	mediator := &ConcreteMediator{}
	A := &ConcreteColleagueA{}
	B := &ConcreteColleagueB{}

	mediator.register(1, A)
	A.mediator = mediator
	mediator.register(2, B)
	B.mediator = mediator

	A.startDialog("How are you?")
	B.startDialog("I am fine. And you?")
}
