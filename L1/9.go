package main

import "fmt"

func main() {
	nums := [5]int{1, 3, 5, 7, 9}

	cNum := make(chan int)
	cNum2 := make(chan int)

	for _, n := range nums {
		go func(n int) {
			cNum <- n
		}(n)

		go func() {
			n2 := <-cNum
			cNum2 <- n2 * n2
		}()

		result := <-cNum2
		fmt.Println(result)
	}
}
