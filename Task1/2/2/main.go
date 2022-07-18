package main

import (
	"fmt"
)

func main() {
	mas := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	for _, elem := range mas {
		go inFile(elem*elem, ch)
		fmt.Println(<-ch)
	}
}

func inFile(elem int, ch chan int) {
	ch <- elem
}
