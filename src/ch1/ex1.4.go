package main

import (
	"bufio"
	"fmt"
	"os"
)

type counter struct {
	count int
	files map[string]int
}

func main() {
	counts := make(map[string]*counter)
	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ex1.4: failed to open %s, err was: %v", file, err)
		}

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			text := scanner.Text()
			if counts[text] == nil {
				counts[text] = &counter{files: make(map[string]int)}
			}

			counts[text].count += 1
			counts[text].files[file] = 1
		}
	}

	for line, counter := range counts {
		if counter.count > 1 {
			fmt.Printf("%s\t%d\t%v\n", line, counter.count, counter.files)
		}
	}
}