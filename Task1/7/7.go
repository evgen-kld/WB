package main

import (
	"fmt"
	"sync"
)

type Map struct {
	mu sync.Mutex
	m  map[int]string
}

func (m *Map) AddNewElem(key int) {
	m.mu.Lock()
	m.m[key] = "elem"
	m.mu.Unlock()
}

func main() {
	m := Map{m: make(map[int]string)}
	for i := 0; i < 100; i++ {
		go m.AddNewElem(i)
	}
	fmt.Println(m.m)
}
