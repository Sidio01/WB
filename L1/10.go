package main

import "fmt"

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -14.5, -32.7, 5.9}
	tempMap := map[int][]float32{}
	for _, num := range temp {
		switch {
		case num <= -30 && num > -40:
			tempMap[-30] = append(tempMap[-30], num)
		case num <= -20 && num > -30:
			tempMap[-20] = append(tempMap[-20], num)
		case num <= -10 && num > -20:
			tempMap[-10] = append(tempMap[-10], num)
		case num >= 0 && num < 10:
			tempMap[0] = append(tempMap[0], num)
		case num >= 10 && num < 20:
			tempMap[10] = append(tempMap[10], num)
		case num >= 20 && num < 30:
			tempMap[20] = append(tempMap[20], num)
		case num >= 30 && num < 40:
			tempMap[30] = append(tempMap[30], num)
		}
	}
	fmt.Println(tempMap)
}
