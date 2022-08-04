package main

import (
	"fmt"
	"sort"
	"strings"
)

func getAnagram(words *[]string) *map[string]*[]string {
	anagrams := make(map[string][]string)
	for _, w := range *words {
		word := strings.ToLower(w)
		wordSorted := sortString(word)
		anagrams[wordSorted] = append(anagrams[wordSorted], word)
	}

	res := make(map[string]*[]string)
	for _, v := range anagrams {
		v := v
		if len(v) > 1 {
			res[v[0]] = &v
			sort.Slice(v, func(i, j int) bool { return v[i] < v[j] })
		}
	}
	return &res
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
	return string(runes)
}

func main() {
	words := &[]string{"пятак", "abc", "тяпка", "пятка", "листок", "слиток", "cab", "abc", "столик"}
	fmt.Println(*words)
	anagrams := getAnagram(words)
	for k, v := range *anagrams {
		fmt.Println(k, *v)
	}

}
