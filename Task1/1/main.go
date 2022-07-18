package main

import "fmt"

//встраиваю структуру action в human
type human struct {
	action
	A int
	B int
}

type action struct {
	A int
	B int
}

func (h *human) plus() {
	h.A++
	h.B++
}

func (a *action) plus() {
	a.A++
	a.B++
}

func main() {
	// инициализирую структуру
	a := human{
		action: action{1, 1},
		A:      0,
		B:      0,
	}
	fmt.Println(a)
	// если названия одинаковые, то будет происходить shadowing (затенение) и будет вызываться метод/переменная высшего уровня
	// если совпадает уровень вложенности и название => коллизия => ошибка
	fmt.Println(a.A)
	fmt.Println(a.action.A)
	a.plus()
	fmt.Println(a)
	a.action.plus()
	fmt.Println(a)
}
