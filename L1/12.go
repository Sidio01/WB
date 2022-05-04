package main

import "fmt"

func main() {
	strings := [5]string{"cat", "cat", "dog", "cat", "tree"}
	result := []string{}

	for _, str := range strings {
		flag := true
		for _, setStr := range result {
			if setStr == str {
				flag = false
			}
		}
		if flag {
			result = append(result, str)
		}
	}
	fmt.Println(result)
}
