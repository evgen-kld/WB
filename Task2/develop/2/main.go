package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	s := strings.Split(`45`, "")
	mas, err := getList(s)
	if err != nil {
		log.Fatalf(err.Error())
	} else {
		fmt.Println(strings.Join(mas, ""))
	}
}

func getList(s []string) ([]string, error) {
	var mas []string
	flag := false
	var char string
	for ind, elem := range s {
		if flag {
			if char == "" {
				char = elem
				continue
			} else if num, err := strconv.Atoi(string(elem)); err == nil {
				for i := 0; i < num; i++ {
					mas = append(mas, elem)
				}
				flag = true
				char = ""
				continue
			}
		}
		if elem == `\` {
			flag = true
			continue
		}
		if num, err := strconv.Atoi(string(elem)); err == nil {
			if ind == 0 {
				return nil, fmt.Errorf("(некорректная строка)")
			}
			for i := 0; i < num-1; i++ {
				mas = append(mas, s[ind-1])
			}
		} else {
			mas = append(mas, elem)
		}
	}
	return mas, nil
}
