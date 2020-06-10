package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func f2c(f float64) float64 {
	return (f - 32) * 5 / 9
}

func main() {
	if len(os.Args[1:]) == 0 {
		log.Fatal("missing temp in fahrenheit")
	}

	src, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatal(err)
	}

	c := f2c(src)

	fmt.Printf("F=%g, C=%g\n", src, c)
}