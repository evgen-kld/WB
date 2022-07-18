package main

import "fmt"

func main() {
	s := "qwerty"
	r := []rune(s)
	fmt.Println(r)
	for left, right := 0, len(r)-1; left < len(r)/2; left, right = left+1, right-1 {
		r[left], r[right] = r[right], r[left]
	}
	fmt.Println(string(r))
}
