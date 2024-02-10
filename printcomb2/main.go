package main

import "github.com/01-edu/z01"

func comb2() {
	for i := ('0' - 48); i <= 98; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(rune(i/10) + '0')
			z01.PrintRune(rune(i%10) + '0')
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10) + '0')
			z01.PrintRune(rune(j%10) + '0')
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}

func main() {
	comb2()
}
