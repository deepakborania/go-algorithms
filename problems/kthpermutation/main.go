package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n, k int) string {

	// blockSize := factorial(n - 1)
	// block := k / blockSize
	numbers := make([]int, n)
	result := ""
	for i := 0; i < n; i++ {
		numbers[i] = i + 1
	}
	k--
	for i := 1; i <= n; i++ {
		f := factorial(len(numbers) - 1)
		idx := k / f
		result += strconv.Itoa(numbers[idx])
		numbers = append(numbers[:idx], numbers[idx+1:]...) // Remove this index from numbers
		k = k % f
	}
	return result
}

func factorial(n int) int {
	result := 1
	for i := n; i > 1; i-- {
		result *= i
	}
	return result
}
func main() {
	fmt.Println(getPermutation(4, 9))
}
