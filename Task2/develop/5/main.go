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

func getContainsList(mas []string, p param, s string) []int {
	newMasElem := make([]int, 0)
	indMas := make([]int, 0)
	unicDict := make(map[int]bool, 0)
	for ind, elem := range mas {
		if p.F {
			if p.i {
				if strings.ToLower(elem) == strings.ToLower(s) {
					indMas = append(indMas, ind)
				}
			} else if elem == s {
				indMas = append(indMas, ind)
			}
		} else if p.i {
			if strings.Contains(strings.ToLower(elem), strings.ToLower(s)) {
				indMas = append(indMas, ind)
			}
		} else if strings.Contains(elem, s) {
			indMas = append(indMas, ind)
		}
	}
	if p.A != 0 {
		for _, i := range indMas {
			for j := 0; j < p.A; j++ {
				unicDict[i+j] = true
			}
		}
	} else if p.B != 0 {
		for _, i := range indMas {
			for j := 0; j <= p.B; j++ {
				unicDict[i-j] = true
			}
		}
	} else if p.C != 0 {
		for _, i := range indMas {
			startInd := i - p.C
			for j := startInd; j <= startInd+p.C*2; j++ {
				unicDict[j] = true
			}
		}
	} else {
		for _, i := range indMas {
			unicDict[i] = true
		}
	}
	for key := range unicDict {
		if key < len(mas) && key >= 0 {
			newMasElem = append(newMasElem, key)
		}
	}
	sort.Ints(newMasElem)
	return newMasElem
}

func printList(masInd []int, mas []string, n bool) {
	fmt.Println(masInd)
	for _, i := range masInd {
		if n {
			fmt.Println(i+1, mas[i])
		} else {
			fmt.Println(mas[i])
		}
	}
}

type param struct {
	A int
	B int
	C int
	c int
	i bool
	v bool
	F bool
	n bool
}

func (p *param) getParams() {
	A := flag.Int("A", 0, "печатать +N строк после совпадения")
	B := flag.Int("B", 0, "печатать +N строк до совпадения")
	C := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	c := flag.Int("c", 0, "количество строк")
	i := flag.Bool("i", false, "игнорировать регистр")
	v := flag.Bool("v", false, "вместо совпадения, исключать")
	F := flag.Bool("F", false, "точное совпадение со строкой")
	n := flag.Bool("n", false, "напечатать номер строки")
	flag.Parse()
	p.A = *A
	p.B = *B
	p.C = *C
	p.c = *c
	p.i = *i
	p.v = *v
	p.F = *F
	p.n = *n
}

func main() {
	var s string
	var filename string
	var in io.Reader
	p := param{}
	p.getParams()
	s = flag.Arg(0)
	if filename = flag.Arg(1); filename != "" {
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
	mas := getList(in)
	fmt.Println(mas)
	masOfInd := getContainsList(mas, p, s)
	printList(masOfInd, mas, p.n)
}
