package main

import "fmt"

func main() {
	mas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	firstChan := make(chan int)
	secondChan := make(chan int)

	go func() {
		for _, elem := range mas {
			firstChan <- elem
		}
		close(firstChan)
	}()

	go func() {
		for {
			x, ok := <-firstChan
			if !ok {
				close(secondChan)
				break
			}
			secondChan <- x * 2
		}
	}()

	for i := range secondChan {
		fmt.Println(i)
	}
}
