package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = map[int]int{}
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m map[int]int, i int, mu *sync.Mutex, wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			m[i] = i * i
			mu.Unlock()
		}(m, i, mu, wg)
	}
	wg.Wait()
	fmt.Println(m)
}
