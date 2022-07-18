package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	mas := []int{2, 4, 6, 8, 10}
	wg.Add(len(mas))
	for _, elem := range mas {
		go inFile(elem * elem)
	}
	wg.Wait()
}

func inFile(elem int) {
	fmt.Println(elem)
	wg.Done()
}
