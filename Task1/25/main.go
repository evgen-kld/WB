package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("123")
	Sleep(5)
	fmt.Println("123")
}

func Sleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}
