package main

import (
	"fmt"
	"strings"
)

func main() {
	s := strings.ToLower("Twer")
	r := []rune(s)
	d := make(map[rune]int)
	for _, val := range r {
		d[val]++
	}
	for i := range d {
		if d[i] != 1 {
			fmt.Println(false)
			return
		}
	}
	fmt.Println(true)
}
