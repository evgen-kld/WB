package main

import "fmt"

func main() {
	mas := []int{5, 4, 8, 10, 99, 25, 54, 65, 45, 1}
	fmt.Println(QuickSort(mas))
}

func QuickSort(mas []int) []int {
	if len(mas) <= 1 {
		return mas
	} else {
		base := mas[0]
		var less []int
		var equal []int
		var greater []int
		for _, i := range mas {
			if i < base {
				less = append(less, i)
			} else if i == base {
				equal = append(equal, i)
			} else {
				greater = append(greater, i)
			}
		}
		return append(append(QuickSort(less), equal...), QuickSort(greater)...)
	}
}
