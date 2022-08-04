package main

import (
	"Task2/pattern/6_factory/pkg"
	"fmt"
)

var types = []string{pkg.PersonalComputerType, pkg.ServerType, "123"}

func main() {
	for _, i := range types {
		computer := pkg.NewComputer(i)
		if computer == nil {
			fmt.Println("Не существует")
			continue
		}
		computer.PrintInfo()
	}
}
