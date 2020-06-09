package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, count := range counts {
		if (count > 1) {
			fmt.Printf("%d\t%s\n", line, count)
		}
	}
}

// to run: go run src/ch1/dup1.go
// enter values and then CTRL+D for the end of line