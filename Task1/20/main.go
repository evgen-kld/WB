package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"
	mas := strings.Fields(s)
	fmt.Println(mas)
	for left, right := 0, len(mas)-1; left < len(mas)/2; left, right = left+1, right-1 {
		mas[left], mas[right] = mas[right], mas[left]
	}
	fmt.Println(mas)
}
