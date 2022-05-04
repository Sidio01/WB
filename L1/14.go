package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := 2
	b := "2"
	c := true
	d := make(chan int)

	aType := reflect.TypeOf(a).Kind()
	bType := reflect.TypeOf(b).Kind()
	cType := reflect.TypeOf(c).Kind()
	dType := reflect.TypeOf(d).Kind()

	if aType == reflect.Int {
		fmt.Println("a is int")
	}
	if bType == reflect.String {
		fmt.Println("b is string")
	}
	if cType == reflect.Bool {
		fmt.Println("c is bool")
	}
	if dType == reflect.Chan {
		fmt.Println("d is channel")
	}
}
