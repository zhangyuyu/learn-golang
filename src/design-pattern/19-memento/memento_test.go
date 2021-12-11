package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	originator := &Originator{
		state: "State 1",
	}
	caretaker := &Caretaker{}
	caretaker.Add(originator.SaveToMemento())

	originator.state = "State 2"
	caretaker.Add(originator.SaveToMemento())

	originator.state = "State 3"
	fmt.Printf("Current.State: %s\n", originator.state)

	originator.RestoreFromMemento(caretaker.Get(0))
	fmt.Printf("First saved State: %s\n", originator.state)

	originator.RestoreFromMemento(caretaker.Get(1))
	fmt.Printf("Second saved State: %s\n", originator.state)
}
