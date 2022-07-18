package main

import "fmt"

func main() {
	m := []int{1, 2, 3, 4, 5}
	i := 3
	m = append(m[:i], m[i+1:]...)
	fmt.Println(m)
}
