package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	sep := []byte("\n")
	counts := make(map[string]int)

	for _, file := range os.Args[1:] {
		f, err := ioutil.ReadFile(file)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: error reading %s: %v", file, err)
			continue
		}

		for _, text := range bytes.Split(f, sep) {
			counts[string(text)]++
		}
	}

	for text, count := range counts {
		if (count > 1) {
			fmt.Printf("%s\t%d\n", text, count)
		}
	}
}