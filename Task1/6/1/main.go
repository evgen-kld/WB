package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	quit := make(chan bool)
	work := make(chan bool)

	wg.Add(1)
	go func(work, quit chan bool) {
		for {
			select {
			case <-work:
				fmt.Println("Горутина работает")
			case <-quit:
				fmt.Println("Горутина завершила работу")
				wg.Done()
				return
			}
		}
	}(work, quit)
	for i := 0; i < 20; i++ {
		work <- true
	}
	quit <- true
	wg.Wait()
}
