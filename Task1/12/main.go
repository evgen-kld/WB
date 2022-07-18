package main

import "fmt"

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(getSet(mas))
}

func getSet(mas []string) []string {
	d := make(map[string]int)
	for _, elem := range mas {
		d[elem] = 0
	}
	var result []string
	for key := range d {
		result = append(result, key)
	}
	return result
}
