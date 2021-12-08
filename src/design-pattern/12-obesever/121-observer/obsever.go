package _21_observer

import "fmt"

type ISubject interface {
	Register(observer IObserver)
	Remove(observer IObserver)
	Notify(msg string)
}

// IObserver 观察者
type IObserver interface {
	Update(msg string)
}

type Subject struct {
	observers []IObserver
}

// Register 注册
func (s *Subject) Register(observer IObserver) {
	s.observers = append(s.observers, observer)
}

// Remove 移除观察者
func (s *Subject) Remove(observer IObserver) {
	for i, ob := range s.observers {
		if ob == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
		}
	}
}

// Notify 通知
func (s *Subject) Notify(msg string) {
	for _, ob := range s.observers {
		ob.Update(msg)
	}
}

type Observer1 struct{}

func (o *Observer1) Update(msg string) {
	fmt.Printf("Observer1: %s\n", msg)
}

type Observer2 struct{}

func (o *Observer2) Update(msg string) {
	fmt.Printf("Observer2: %s\n", msg)
}
