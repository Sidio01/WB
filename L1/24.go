package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func distance(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

func main() {
	p1 := Point{x: 1, y: 1}
	p2 := Point{x: 4, y: 5}
	fmt.Println(distance(p1, p2))
}
