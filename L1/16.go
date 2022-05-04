package main

import "fmt"

func quickSort(m []int) []int {
	if len(m) < 2 {
		return m
	}
	mid := len(m) / 2
	left := []int{}
	right := []int{}
	for idx := range m {
		if m[idx] <= m[mid] && idx != mid {
			left = append(left, m[idx])
		} else if m[idx] > m[mid] {
			right = append(right, m[idx])
		}
	}

	l := quickSort(left)
	r := quickSort(right)

	result := []int{}
	result = append(result, l...)
	result = append(result, m[mid])
	result = append(result, r...)

	return result
}

func main() {
	m := []int{4, 6, 1, 3, 9, -2, 11, 7, 9}
	fmt.Println(quickSort(m))
}
