package main

import "fmt"

func main() {
	// Вариант 1
	a := 10
	b := 8
	fmt.Printf("a - %v, b - %v\n", a, b)
	b = a + b
	a = b - a
	b = b - a
	fmt.Printf("a - %v, b - %v\n", a, b)

	// Вариант 2
	a = 10
	b = 8
	fmt.Printf("a - %v, b - %v\n", a, b)
	a, b = b, a
	fmt.Printf("a - %v, b - %v\n", a, b)
}
