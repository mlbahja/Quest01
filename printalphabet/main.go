package main

import "github.com/01-edu/z01"

func main() {
	for i := 'a'; i <= 'z'; i = i + 1 {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
