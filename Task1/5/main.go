package main

import (
	"fmt"
	"time"
)

func Write(c chan string) {
	fmt.Println(<-c)
}

func main() {
	var t int
	fmt.Scan(&t)
	c := make(chan string)
	go Write(c)

	for {
		c <- "123"
	}

	time.Sleep(time.Duration(t) * time.Second)
}
