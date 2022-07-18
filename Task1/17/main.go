package main

import (
	"fmt"
	"sort"
)

func main() {
	mas := []int{5, 4, 8, 10, 99, 25, 54, 65, 45, 1}
	sort.Ints(mas)
	fmt.Println(mas)
	if elem, ok := binarySearch(mas, -5); ok {
		fmt.Println(elem)
	} else {
		fmt.Println("Не существует")
	}
}

func binarySearch(mas []int, item int) (int, bool) {
	low := 0
	high := len(mas) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := mas[mid]
		if guess == item {
			return mid, true
		} else if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, false
}
