package main

import (
	"flag"
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func main() {
	var host string
	flag.StringVar(&host, "h", "time.apple.com", "set host address")
	flag.Parse()
	tm, err := GetExactTime(host)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tm.String())
}

func GetExactTime(host string) (*time.Time, error) {
	tm, ok := ntp.Time(host)
	if ok != nil {
		return nil, ok
	}
	return &tm, nil
}
