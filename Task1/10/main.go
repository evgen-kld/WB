package main

import (
	"fmt"
)

func main() {
	var mas = []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -9.9, -15, -10.1, -10, 5}
	di := make(map[int][]float64)
	for _, i := range mas {
		d := int(i/10) * 10
		if i < 0 {
			d = d - 10
		}
		di[d] = append(di[d], i)
	}
	fmt.Println(di)
}
