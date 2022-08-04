package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	data := []string{"Стратегия", "—", "это", "поведенческий", "паттерн", "проектирования,", "который", "определяет", "семейство", "схожих", "алгоритмов", "и", "помещает"}
	ApplyingStrategy(data, Upper{})
	ApplyingStrategy(data, SortLength{})
	ApplyingStrategy(data, СaesarСipher{-12})
	ApplyingStrategy(data, SortLexicographic{})
}

func ApplyingStrategy(s []string, operation IOperation) {
	copyData := make([]string, len(s))
	copy(copyData, s)
	fmt.Printf("%-23s", fmt.Sprintf("%v", reflect.TypeOf(operation)))
	fmt.Print("Len:", len(copyData), " ")
	operation.Processing(copyData)
	fmt.Println(copyData)
}

type IOperation interface{ Processing([]string) }

type Upper struct{}

type SortLength struct{}

type СaesarСipher struct{ level int }

type SortLexicographic struct{}

func (Upper) Processing(s []string) {
	size := len(s)
	for i := 0; i < size; i++ {
		s[i] = strings.ToUpper(s[i])
	}
}

func (SortLength) Processing(s []string) {
	size, swapper := len(s), reflect.Swapper(s)

	for i := 0; i < size; i++ {
		for x := i + 1; x < size; x++ {
			if len(s[i]) > len(s[x]) {
				swapper(x, i)
			}
		}
	}
}

func (SortLexicographic) Processing(s []string) {
	size, swapper := len(s), reflect.Swapper(s)
	compare := func(a, b string) bool {
		sa, sb := len(a), len(b)
		for i := 0; i < sa && i < sb; i++ {
			if a[i] == b[i] {
				continue
			}
			return a[i] < b[i]
		}
		return sa < sb
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if compare(s[i], s[j]) {
				swapper(i, j)
			}
		}
	}
}

func (c СaesarСipher) Processing(s []string) {
	size := len(s)
	up := func(s string) string {
		res := make([]rune, len(s))
		for i, v := range s {
			res[i] = v + rune(c.level)
		}
		return string(res)
	}
	for i := 0; i < size; i++ {
		s[i] = up(s[i])
	}
}
