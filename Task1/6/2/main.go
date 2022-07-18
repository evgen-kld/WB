package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan bool)
	
	wg.Add(1)

	go func(ch chan bool) {
		for {
			_, ok := <-ch
			if !ok {
				wg.Done()
				fmt.Println("Завершение")
				return
			}
			fmt.Println("Работа")
		}
	}(ch)

	for i := 0; i < 20; i++ {
		ch <- true
	}

	close(ch)
	wg.Wait()
}
