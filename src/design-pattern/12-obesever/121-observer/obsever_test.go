package _21_observer

import "testing"

func TestSubject_Notify(t *testing.T) {
	subject := &Subject{}
	subject.Register(&Observer1{})
	subject.Register(&Observer2{})

	subject.Notify("Hi")
}
