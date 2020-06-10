package main

import (
	"fmt"
	"io"
	"time"
	"io/ioutil"
	"net/http"
	"os"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetchall: err while reading response for %s. err was: %v", url, err)
		return
	}

	err = resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetchall: err while closing the connection for %s. err was: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, bytes, url)
}

func main() {
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}