package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "snow dog sun"
	buf := strings.Split(text, " ")
	result := []string{}
	for i := len(buf) - 1; i >= 0; i-- {
		result = append(result, buf[i])
	}
	fmt.Println(strings.Join(result, " "))
}
