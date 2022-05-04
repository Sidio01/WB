package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	Count int32
}

func (c *Counter) inc() {
	atomic.AddInt32(&c.Count, 1)
}

func main() {
	var counter = Counter{Count: 0}
	var mu sync.Mutex
	wg := &sync.WaitGroup{}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			counter.inc()
			mu.Unlock()
		}(wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
