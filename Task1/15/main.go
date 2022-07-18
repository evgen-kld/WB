package main

import (
	"fmt"
	"strings"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100]
	fmt.Println(justString)
	rune := make([]rune, 100)
	i := 0
	for _, r := range v {
		if i >= 100 {
			break
		}
		rune[i] = r
		i++
	}
	fmt.Println(string(rune))
}

func createHugeString(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString("럼")
		//builder.WriteString("i")
	}
	return builder.String()
}

func main() {
	someFunc()
}
