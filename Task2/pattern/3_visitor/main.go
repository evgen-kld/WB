package main

import "fmt"

type shape interface {
	getType() string
	getABC() string
	accept(visitor)
}

type square struct {
	side int
}

//func (s *square) accept(v visitor) {
//	v.visitForSquare(s)
//}

func (s *square) getType() string {
	return "Square"
}

func (s *square) GetABC() string {
	return "ABC"
}

type visitor interface {
	visitForSquare(*square)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for square")
}

//type middleCoordinates struct {
//	x int
//	y int
//}

//func (a *middleCoordinates) visitForSquare(s *square) {
//	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
//	fmt.Println("Calculating middle point coordinates for square")
//}

func main() {
	square := &square{side: 2}
	fmt.Println(square.getType())
	fmt.Println(square.GetABC())
	//areaCalculator := &areaCalculator{}
	//square.accept(areaCalculator)

	fmt.Println()

	//middleCoordinates := &middleCoordinates{}
	//square.accept(middleCoordinates)
}
