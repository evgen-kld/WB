package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Worker(ch chan int, id int) {
	for {
		fmt.Println("Worker", id, "значение:", <-ch)
	}
}

func main() {
	ch := make(chan int)
	sig := make(chan os.Signal)

	signal.Notify(sig, syscall.SIGINT)

	var workerCount int
	fmt.Scan(&workerCount)

	for i := 1; i <= workerCount; i++ {
		go Worker(ch, i)
	}

	var res int
	for {
		select {
		case ch <- res:
			res++
		case <-sig:
			fmt.Println("Завершение программы")
			return
		}
	}
}
