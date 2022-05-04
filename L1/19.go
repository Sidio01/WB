package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	stringToReverse := "главрыба"
	reversedString := make([]string, utf8.RuneCountInString(stringToReverse))
	i := utf8.RuneCountInString(stringToReverse) - 1
	for _, r := range stringToReverse {
		reversedString[i] = string(r)
		i--
	}
	fmt.Println(strings.Join(reversedString, ""))
}
