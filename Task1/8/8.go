package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64
	fmt.Scan(&n)
	bin := strconv.FormatInt(n, 2)
	fmt.Println(bin)
	var i int
	fmt.Scan(&i)
	bin = bin[:i] + "0" + bin[i+1:]
	fmt.Println(bin)
	n, _ = strconv.ParseInt(bin, 2, 64)
	fmt.Println(n)
}
