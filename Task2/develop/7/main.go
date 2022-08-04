package main

import (
	"fmt"
	"time"
)

func or(channels ...<-chan interface{}) <-chan interface{} {
	single := make(chan interface{})
	done := make(chan struct{})

	for i := range channels {
		go func(ch <-chan interface{}) {
			select {
			case tmp := <-ch:
				close(single)
				single <- tmp
			case <-done:
				return
			}
		}(channels[i])
	}

	<-done

	return single
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf(`fone after %v`, time.Since(start))

}
