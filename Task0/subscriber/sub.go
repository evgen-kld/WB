package main

import (
	"Task0/JSON"
	"Task0/cache"
	"Task0/db"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"sync"
)

func main() {

	s := JSON.StructJSON{}

	sc, _ := stan.Connect("test-cluster", "sub")
	defer sc.Close()
	sc.Subscribe("json_channel", func(m *stan.Msg) {
		if err := json.Unmarshal(m.Data, &s); err != nil {
			fmt.Println(err)
			return
		}
		str, ok := cache.AddToCache(s)
		fmt.Println(str)
		if ok {
			db.Insert(s)
		}
	})
	Block()
}

func Block() {
	w := sync.WaitGroup{}
	w.Add(1)
	w.Wait()
}
