package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}

		res, err := http.Get(url)
		if (err != nil) {
			fmt.Fprintf(os.Stderr, "fetch: failed to fetch %s, err was: %v\n", url, err)
			continue
		}

		if _, err2 := io.Copy(os.Stdout, res.Body); err2 != nil {
			fmt.Fprintf(os.Stderr, "fetch: failed to read %s, err was: %v\n", url, err2)
		}
	}
}