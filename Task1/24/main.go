package main

import (
	"fmt"
	"math"
)

type Pointer struct {
	x int
	y int
}

func getDistance(a, b Pointer) float64 {
	return math.Pow(math.Pow(float64(a.x+b.x), 2)+math.Pow(float64(a.y+b.y), 2), 0.5)
}

func NewPointer(x, y int) Pointer {
	return Pointer{
		x: x,
		y: y,
	}
}

func main() {
	a := NewPointer(8, 2)
	b := NewPointer(4, 2)
	fmt.Println(getDistance(a, b))
}
