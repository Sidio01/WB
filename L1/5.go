package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	workTime := 5 * time.Second
	ctx, _ := context.WithTimeout(context.Background(), workTime)

	c := make(chan int)

	go func(ctx context.Context, in chan int) {
		i := 0
		for {
			in <- i
			i++
		}
	}(ctx, c)

LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case n := <-c:
			fmt.Println(n)
		}

	}
}
