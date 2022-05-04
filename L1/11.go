package main

import "fmt"

func main() {
	n := []int{4, 2, 3, 1}
	m := []int{5, 4, 6, 3}
	intersection := []int{}

	for _, i := range n {
		for _, j := range m {
			if i == j {
				intersection = append(intersection, i)
			}
		}
	}
	fmt.Println(intersection)
}
