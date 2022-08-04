package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"path"
)

func download(url, filename string) (err error) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	f, _ := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return
}

func main() {
	pUrl := flag.String("url", "https://siongui.github.io/index.html", "URL to be processed")
	flag.Parse()
	url := *pUrl
	filename := path.Base(url)
	os.Stat(filename)
	download(url, filename)
}
