package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "));
}

// to run: go run src/ch1/ex1.1.go foo bar