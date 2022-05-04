package main

import "fmt"

func createSlice(l int) []int {
	sl := make([]int, l)
	for i := 0; i < len(sl); i++ {
		sl[i] = i + 1
	}
	return sl
}

func binSearch(slice []int, target int, step int) ([]int, int) {
	middle := len(slice) / 2
	if slice[middle] == target {
		return []int{target}, step + 1
	} else if target < slice[middle] {
		return binSearch(slice[:middle], target, step+1)
	} else {
		return binSearch(slice[middle:], target, step+1)
	}
}

func main() {
	slice := createSlice(100)
	targetNum := 42
	step := 1

	fmt.Println(binSearch(slice, targetNum, step))
}
