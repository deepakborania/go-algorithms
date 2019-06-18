package main

import "fmt"

/*
 Given a number array to represent different coin denominations and a total amount ‘T’,
 we need to find all the different ways to make a change for ‘T’ with the given coin denominations.
 We can assume an infinite supply of coins, therefore, each coin can be chosen multiple times
*/

func coinChange(denoms []int, total int) int {
	dp := make([][]int, len(denoms))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, total+1)
	}
	return helper(denoms, total, 0, dp)
}

func helper(denoms []int, total int, curIdx int, dp [][]int) int {
	if len(denoms) == 0 || curIdx >= len(denoms) {
		return 0
	}
	if total == 0 {
		return 1
	}
	if dp[curIdx][total] != 0 {
		return dp[curIdx][total]
	}
	sum1 := 0
	if denoms[curIdx] <= total {
		sum1 = helper(denoms, total-denoms[curIdx], curIdx, dp)
	}
	sum2 := helper(denoms, total, curIdx+1, dp)
	dp[curIdx][total] = sum1 + sum2
	return dp[curIdx][total]
}

func coinChangeBU(denoms []int, total int) int {
	if len(denoms) == 0 || total == 0 {
		return 0
	}
	N := len(denoms)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, total+1)
	}
	for i := 0; i < N; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < N; i++ {
		for t := 1; t <= total; t++ {
			if i > 0 {
				dp[i][t] = dp[i-1][t]
			}
			if denoms[i] <= t {
				dp[i][t] += dp[i][t-denoms[i]]
			}
		}
	}
	return dp[N-1][total]
}
func main() {
	denoms := []int{1, 2, 3, 5, 7, 11}
	total := 123
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", coinChange(denoms, total))
	fmt.Println("Denominations=", denoms, "Total Required =", total, "Number of ways=", coinChangeBU(denoms, total))
}
