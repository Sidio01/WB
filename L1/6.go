package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func contextCancel(c context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-c.Done():
			fmt.Println("goroutine stop by cancel")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}

func contextTimeout(c context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-c.Done():
			fmt.Println("goroutine stop by timeout")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}

func cancelByChannel(c chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-c:
			fmt.Println("goroutine stop by channel")
			time.Sleep(time.Millisecond * 10)
			return
		}
	}
}
func main() {
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	t := time.Second * time.Duration(5)

	ctxCancel, cancel := context.WithCancel(ctx)
	ctxTimeout, _ := context.WithTimeout(ctx, t)
	ch := make(chan struct{})

	fmt.Println("goroutine start by cancel")
	wg.Add(1)
	go contextCancel(ctxCancel, wg)
	fmt.Println("goroutine start by timeout")
	wg.Add(1)
	go contextTimeout(ctxTimeout, wg)
	fmt.Println("goroutine start by channel")
	wg.Add(1)
	go cancelByChannel(ch, wg)

	select {
	case <-time.After(t):
		cancel()
		ch <- struct{}{}
	}
	wg.Wait()
}
