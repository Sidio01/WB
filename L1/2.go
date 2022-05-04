package main

import (
	"fmt"
	"sync"
)

func main() {
	m := [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	for _, i := range m {
		wg.Add(1)
		go func(x int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(x * x)
		}(i, wg)
	}
	wg.Wait()
}
