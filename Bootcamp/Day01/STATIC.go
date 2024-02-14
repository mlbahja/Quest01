package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nq := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nq[0])
	q, _ := strconv.Atoi(nq[1])

	scanner.Scan()
	arrayStr := strings.Fields(scanner.Text())
	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i], _ = strconv.Atoi(arrayStr[i])
	}

	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + array[i-1]
	}
	for i := 0; i < q; i++ {
		scanner.Scan()
		query := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(query[0])
		b, _ := strconv.Atoi(query[1])+
		sum := prefixSum[b] - prefixSum[a-1]
		fmt.Println(sum)
	}
}
