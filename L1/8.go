package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int64 = 43 // 000000000000000000000000000000000000000000000000000000000101011
	var bit int64 = 1  // 0 or 1
	var bitPosition int64 = 0
	n, _ := strconv.ParseInt("111111111111111111111111111111111111111111111111111111111111111", 10, 64) // 9223372036854775807
	if bit == 0 {
		fmt.Println(strconv.FormatInt(num&(n-1<<bitPosition), 2))
	} else {
		fmt.Println(strconv.FormatInt(num|(1<<bitPosition), 2))
	}
}
