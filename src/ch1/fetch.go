package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "fetch: failed to fetch %s, err was: %v\n", url, err)
			continue
		}

		b, err2 := ioutil.ReadAll(res.Body)
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "fetch: failed to read %s, err was: %v\n", url, err2)
		}
		res.Body.Close()

		fmt.Printf("content of %s is: %s\n", url, b)
	}
}