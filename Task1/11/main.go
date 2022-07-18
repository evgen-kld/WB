package main

import "fmt"

func Intersection(a, b []int) []int {
	m := make(map[int]int)
	var result []int
	for _, el := range a {
		if _, ok := m[el]; ok == false {
			m[el] = 1
		} else {
			m[el] += 1
		}
	}
	for _, el := range b {
		if count, ok := m[el]; ok && count > 0 {
			m[el] -= 1
			result = append(result, el)
		}
	}
	return result
}

func main() {
	a := []int{3, 5, 7, 8, 15}
	b := []int{15, 5, 8, 0, -1, 2, 4, 23}
	fmt.Println(Intersection(a, b))
}
