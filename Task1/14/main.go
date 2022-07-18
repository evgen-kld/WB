package main

import "fmt"

func main() {
	var a interface{} = make(chan int)
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel")
	default:
		fmt.Println("Unknown")
	}
}
