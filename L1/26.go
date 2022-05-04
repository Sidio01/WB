package main

import "fmt"

func uniqString(s string) bool {
	buf := []rune{}
	for _, r := range s {
		for _, br := range buf {
			if r == br {
				return false
			}
		}
		buf = append(buf, r)
	}
	return true
}

func main() {
	fmt.Println(uniqString("abcd"))
	fmt.Println(uniqString("abCdefAaf"))
	fmt.Println(uniqString("aabcd"))
}
