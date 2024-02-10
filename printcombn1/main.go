package main

import (
	"fmt"
)

func printComb(n int) {
	if n < 1 || n > 9 {
		return
	}
	combinations := make([]int, n)
	for {
		valid := true
		for i := 0; i < n-1; i++ {
			if combinations[i] > combinations[i+1] {
				valid = false
				break
			}
		}
		if valid {
			for _, num := range combinations {
				fmt.Printf("%d", num)
			}
			fmt.Print(", ")
		}
		carry := 1
		for i := n - 1; i >= 0; i-- {
			combinations[i] += carry
			if combinations[i] > 9 {
				combinations[i] = 0
				carry = 1
			} else {
				carry = 0
				break
			}
		}
		if carry != 0 {
			break
		}
	}
}

func main() {
	printComb(9)
}
