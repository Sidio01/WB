package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	elementToDelete := 1
	x = append(x[:elementToDelete-1], x[elementToDelete:]...)
	fmt.Println(x)
}
