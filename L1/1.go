package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	Human
}

func (h Human) sayHello() {
	fmt.Println("Hello!")
}

func main() {
	var h = Human{name: "Ivan"}
	h.sayHello()
	var a = Action{}
	a.sayHello()
}
