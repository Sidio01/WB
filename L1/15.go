package main

import (
	"fmt"
	"math/rand"
)

func createHugeString(size int) string {
	letters := []string{"а", "б", "в", "г", "д", "е"}
	var v string
	for i := 0; i < size; i++ {
		v += letters[rand.Intn(5)]
	}
	return v
}

func someFunc() (string, string) {
	// Если в строке будут присутствовать символы unicode, то такой подход не сработает, так как слайс по строке возвращает слайс байт.
	// В unicode символ может занимать более 1 байта.
	v := createHugeString(1 << 10)
	justString := v[:100]
	return v, justString
}

func main() {
	v, justString := someFunc()

	correctV := []rune(createHugeString(1 << 10))
	correctJustString := correctV[:100]

	fmt.Printf("v = %v\n", v)
	fmt.Printf("justString = %v\n\n", justString)

	fmt.Printf("correctV = %v\n", string(correctV))
	fmt.Printf("correctJustString = %v\n", string(correctJustString))
}
