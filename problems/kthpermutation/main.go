package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n, k int) string {

	blockSize := factorial(n - 1)
	block := k / blockSize
	return strconv.Itoa(block)
}

func factorial(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}
func main() {
	fmt.Println(getPermutation(4, 14))
}
