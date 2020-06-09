package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(input *os.File, counts map[string]int) {
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		text := scanner.Text()
		counts[text]++
	}
}

func main() {
	counts := make(map[string]int)

	if len(os.Args[1:]) == 0 {
		countLines(os.Stdin, counts)
	} else {
		files := os.Args[1:]

		for _, file := range files {
			f, err := os.Open(file)

			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLines(f, counts)
			f.Close()
		}
	}

	for text, count := range counts {
		if count > 1 {
			fmt.Printf("%s\t:%d\n", text, count)
		}
	}
}