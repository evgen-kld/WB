package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func getList(in io.Reader) []string {
	buf := bufio.NewScanner(in)
	s := make([]string, 0)
	for buf.Scan() {
		s = append(s, buf.Text())
	}
	return s
}

func uniqueSort(mas []string) []string {
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

func rewriteFile(mas []string, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	for _, i := range mas {
		_, err = f.WriteString(i + "\n")
		if err != nil {
			panic(err)
		}
	}
}

func reverseSort(mas []string) []string {
	for i, j := 0, len(mas)-1; i < j; i, j = i+1, j-1 {
		mas[i], mas[j] = mas[j], mas[i]
	}
	return mas
}

func columnSort(mas []string, k int) []string {
	l := make([][]string, 1)
	for ind, elem := range mas {
		l = append(l, make([]string, 0))
		s_list := strings.Split(elem, " ")
		l[ind] = append(l[ind], s_list...)
	}
	l = l[:len(l)-1]
	sort.SliceStable(l, func(i, j int) bool {
		return l[i][k-1] < l[j][k-1]
	})
	new_mas := make([]string, 0)
	for _, elem := range l {
		new_mas = append(new_mas, strings.Join(elem, " "))
	}
	return new_mas
}

func main() {
	var k int
	var n bool
	var r bool
	var u bool
	var in io.Reader
	var filename string

	flag.IntVar(&k, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.Parse()

	if filename = flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("Указанного файла не существует")
			return
		}
		defer f.Close()
		fmt.Println("Успешно")
		in = f
	} else {
		fmt.Println("Ошибка чтения файла")
		return
	}
	fmt.Println(k)
	mas := getList(in)

	mas = uniqueSort(mas)
	fmt.Println("123")

	sort.Strings(mas)

	if k != 0 {
		mas = columnSort(mas, k)
	}

	if r {
		mas = reverseSort(mas)
	}

	fmt.Println(mas)
	rewriteFile(mas, filename)

}
