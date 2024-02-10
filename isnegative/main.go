package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func isnegative(n int) {
	//F >>> [0,1[
	if n >= 1 || n == 0 {
		z01.PrintRune('F')
	} else {
		z01.PrintRune('T')
	}
}
func main() {
	isnegative(1)
	fmt.Println()
	isnegative(0)
	fmt.Println()
	isnegative(-1)
	fmt.Println()
	isnegative(3)
}
