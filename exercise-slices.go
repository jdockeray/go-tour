// Exercise: Slices
package main

import (
	"fmt"

	"golang.org/x/tour/pic"
)

// Pic ...
// Implement Pic. It should return a slice of length dy,
// each element of which is a slice of dx 8-bit unsigned
// integers. When you run the program, it will display
// your picture, interpreting the integers as grayscale
// (well, bluescale) values.
func Pic(dx, dy int) [][]uint8 {
	sliceA := make([][]uint8, dx)

	for i := 0; i < dx; i++ {
		sliceA[i] = make([]uint8, dy)
	}
	fmt.Println(sliceA)

	return sliceA
}

func main() {
	pic.Show(Pic)
}
