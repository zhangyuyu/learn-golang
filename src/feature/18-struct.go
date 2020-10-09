package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{name: "Amy", age: 18})
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Calvin"})

	fmt.Println(&person{name: "David"})
	fmt.Println(newPerson("Fred"))

	s := person{name: "John", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 55
	fmt.Println(sp.age)
	fmt.Println(s.age)
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 45
	return &p
}
