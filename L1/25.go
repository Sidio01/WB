package main

import (
	"fmt"
	"time"
)

func goSleep(x int, t time.Duration) {
	<-time.After(t * time.Duration(x))
}

func main() {
	t := 2
	period := "milliseconds"
	fmt.Printf("Sleep for %v %v\n", t, period)
	goSleep(2, time.Millisecond)
	fmt.Printf("Awake after %v %v", t, period)
}
