package main

import "fmt"

func main() {
	x := 8
	y := 12

	for y != 0 {
		x, y = y, x%y
	}

	fmt.Println(x)
}