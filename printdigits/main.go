package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i = i + 1 {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
