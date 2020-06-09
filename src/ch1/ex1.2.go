package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Println("index:", i, "value:", arg);
	}
}

// to run:  go run src/ch1/ex1.2.go foo bar