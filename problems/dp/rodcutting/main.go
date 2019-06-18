package main

import (
	"fmt"
	"strconv"

	"github.com/deepakborania/go-algorithms/utils"
)

// Given a rod of length ‘n’, we are asked to cut the rod and sell the pieces in a way that will
// maximize the profit. We are also given the price of every piece of length 'i' where '1 <= i <= n'.
func rodCuttingProfit(lengths, prices []int, n int) int {
	return helper(lengths, prices, n, 0)
}

func helper(lengths, prices []int, remLen, curIdx int) int {
	if len(lengths) == 0 || len(prices) != len(lengths) || curIdx >= len(lengths) || remLen == 0 {
		return 0
	}
	profit1 := 0
	if lengths[curIdx] <= remLen {
		profit1 = prices[curIdx] + helper(lengths, prices, remLen-lengths[curIdx], curIdx)
	}
	profit2 := helper(lengths, prices, remLen, curIdx+1)
	return utils.Max(profit1, profit2)
}

func main() {
	lengths := []int{1, 2, 3, 4, 5}
	prices := []int{2, 6, 7, 10, 13}
	fmt.Println("Lengts=", lengths, "Prices=", prices, "Length=", strconv.Itoa(5), "Profit=", rodCuttingProfit(lengths, prices, 5))
}
