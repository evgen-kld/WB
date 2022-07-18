package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	count int
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	wg.Done()
}

func (c *Counter) GetValue() int {
	return c.count
}

func main() {
	wg := sync.WaitGroup{}
	c := Counter{
		mu:    sync.Mutex{},
		count: 0,
	}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go c.Increment(&wg)
	}
	wg.Wait()
	fmt.Println(c.GetValue())
}
