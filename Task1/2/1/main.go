package main

import (
	"fmt"
	"time"
)

func main() {
	mas := []int{2, 4, 6, 8, 10}
	for _, elem := range mas {
		go inFile(elem * elem)
	}
	time.Sleep(1 * time.Second)
}

func inFile(elem int) {
	fmt.Println(elem)
}
