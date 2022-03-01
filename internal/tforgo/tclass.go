package tforgo

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Say() {
	fmt.Println("Hello, my name is", p.name, " age is ", p.age)
}

func Tclass() {
	p1 := Person{"Tom", 18}
	p1.Say()
}
