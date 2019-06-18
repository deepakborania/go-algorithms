package main

import (
	"fmt"
	"math"

	"github.com/deepakborania/go-algorithms/utils"
)

func minCoinChange(denoms []int, total int) int {
	dp := make([][]int, len(denoms))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, total+1)
	}
	result := helper(denoms, total, 0, dp)
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func helper(denoms []int, total int, idx int, dp [][]int) int {
	if total == 0 {
		return 0
	}
	if len(denoms) == 0 || idx >= len(denoms) {
		return math.MaxInt64
	}
	if dp[idx][total] != 0 {
		return dp[idx][total]
	}
	c1 := math.MaxInt64
	if denoms[idx] <= total {
		res := helper(denoms, total-denoms[idx], idx, dp)
		if res != math.MaxInt64 {
			c1 = res + 1
		}
	}
	c2 := helper(denoms, total, idx+1, dp)
	res := utils.Min(c1, c2)
	dp[idx][total] = res
	return dp[idx][total]
}

func minCoinChangeBU(denoms []int, total int) int {
	N := len(denoms)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, total+1)
	}
	for i := 0; i < N; i++ {
		for j := 1; j <= total; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for i := 0; i < N; i++ {
		for t := 1; t <= total; t++ {
			if i > 0 {
				dp[i][t] = dp[i-1][t]
			}
			if denoms[i] <= t && dp[i][t-denoms[i]] != math.MaxInt64 {
				dp[i][t] = utils.Min(dp[i][t], dp[i][t-denoms[i]]+1)
			}
		}
	}
	return dp[N-1][total]
}
func main() {
	denoms := []int{1, 2, 3, 5, 7, 11}
	total := 5
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", minCoinChange(denoms, total))
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", minCoinChangeBU(denoms, total))
	total = 1123
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", minCoinChange(denoms, total))
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", minCoinChangeBU(denoms, total))
	total = 7
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", minCoinChange(denoms, total))
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", minCoinChangeBU(denoms, total))
}
