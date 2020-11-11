package main

import (
	"flag"
	"fmt"
)

/*
This be a number, polygonal
From (s, n) to return, from triangular to hexagonal.
A simple formula is all you need.
So compile this program, give s and n a number, and then you should see.
*/

func getPolygonal(s, n int) int {
	polygonal := (s-2)*((n*(n-1))/2) + n
	return polygonal
}

func main() {
	s := flag.Int("s", 3, "shape of the number")
	n := flag.Int("n", 4, "number index")
	flag.Parse()
	sPointer := *s
	nPointer := *n
	result := getPolygonal(sPointer, nPointer)
	fmt.Printf("%d", result)
}
