package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var s, sep string

	slowStart := time.Now()
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	slowEnd := time.Now()
	slowDuration := slowEnd.Sub(slowStart)

	fastStart := time.Now()
	s2 := strings.Join(os.Args[1:], " ")
	fastEnd := time.Now()
	fastDuration := fastEnd.Sub(fastStart)

	fmt.Println("slow one:", s, slowDuration)
	fmt.Println("fast one", s2, fastDuration)
}

// to run: go run src/ch1/ex1.3.go foo bar