package main

import "fmt"

type usb interface {
	connectWithUsbCable()
}

type memoryCard struct{}

func (c memoryCard) insert() {
	fmt.Println("Карта памяти успешно вставлена!")
}

func (c memoryCard) copy() {
	fmt.Println("Данные скопированы на компьютер!")
}

type adapter struct {
	memoryCard
}

func (a adapter) connectWithUsbCable() {
	a.insert()
	a.copy()
}

func main() {
	card := memoryCard{}
	a := adapter{card}
	a.connectWithUsbCable()
}
