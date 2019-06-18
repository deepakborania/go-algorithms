package main

import (
	"fmt"

	"github.com/deepakborania/go-algorithms/utils"
)

func solveKnapsack(profits, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}
	return skr(profits, weights, capacity, 0, dp)
}
func skr(profits, weights []int, capacity, curIdx int, dp [][]int) int {
	if len(profits) == 0 || len(profits) != len(weights) || capacity == 0 || curIdx >= len(profits) {
		return 0
	}
	if dp[curIdx][capacity] != 0 {
		return dp[curIdx][capacity]
	}
	profit1 := 0
	if weights[curIdx] <= capacity {
		profit1 = profits[curIdx] + skr(profits, weights, capacity-weights[curIdx], curIdx, dp)
	}
	profit2 := skr(profits, weights, capacity, curIdx+1, dp)
	result := utils.Max(profit1, profit2)
	dp[curIdx][capacity] = result
	return result
}

func solveKnapsackBU(profits, weights []int, capacity int) int {
	if len(profits) == 0 || len(profits) != len(weights) || capacity == 0 {
		return 0
	}
	N := len(profits)
	dp := make([][]int, N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}
	for i := 0; i < N; i++ {
		for c := 1; c <= capacity; c++ {
			profit1, profit2 := 0, 0
			if weights[i] <= c {
				profit1 = profits[i] + dp[i][c-weights[i]]
			}
			if i > 0 {
				profit2 = dp[i-1][c]
			}
			dp[i][c] = utils.Max(profit1, profit2)
		}

	}
	return dp[N-1][capacity]
}

// func skr()
func main() {
	p := []int{15, 50, 90, 60, 4, 34, 12, 13, 23, 44, 12, 45, 22, 54, 23, 10, 20, 30, 50, 30, 22, 11, 21, 25, 36, 17, 38, 27}
	fmt.Println(len(p))
	w := []int{1, 3, 5, 4, 9, 11, 9, 2, 4, 3, 7, 9, 4, 8, 7, 2, 7, 9, 4, 2, 5, 8, 3, 6, 11, 7, 13, 6}
	fmt.Println(len(w))
	c := 7800
	fmt.Println("Profits=", p, "Weights=", w, "Capacity=", c, "Maximum Profit", solveKnapsack(p, w, c))
	fmt.Println("Profits=", p, "Weights=", w, "Capacity=", c, "Maximum Profit", solveKnapsackBU(p, w, c))
}
