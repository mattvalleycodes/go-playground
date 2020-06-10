package main

import "fmt"

func main() {
	n := 5
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	fmt.Println(x)
}