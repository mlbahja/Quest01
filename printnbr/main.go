package main

import (
	"fmt"
	"math"

	"github.com/01-edu/z01"
)

func printnbr(n int) {
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		printnbr(223372036854775808)
		return
	}
	if n <= 9 {
		z01.PrintRune(rune(n + '0'))
	}
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}
	if n > 9 {
		printnbr(n / 10)
		printnbr(n % 10)
	}
}
func main() {
	fmt.Println(math.MinInt)
	printnbr(-26598)
	fmt.Println()
	printnbr(0)
	fmt.Println()
	printnbr(36987)
	fmt.Println()

}
