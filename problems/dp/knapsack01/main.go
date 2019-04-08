package main

import (
	"fmt"
	"time"

	"github.com/deepakborania/go-algorithms/utils"
)

func SolveKnapsackRecursive(profits, weights []int, capacity int) int {
	return skr(profits, weights, capacity, 0)
}

func skr(profits, weights []int, capacity, curIdx int) int {
	if curIdx >= len(profits) || capacity <= 0 {
		return 0
	}
	p1 := 0
	if weights[curIdx] <= capacity {
		p1 = profits[curIdx] + skr(profits, weights, capacity-weights[curIdx], curIdx+1)
	}
	p2 := skr(profits, weights, capacity, curIdx+1)
	return utils.Max(p1, p2)
}

func SolveKnapsackMemoized(profits, weights []int, capacity int) int {
	dp := make([][]int, len(profits))
	for i := 0; i < len(profits); i++ {
		dp[i] = make([]int, capacity+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}
	return skrMemoized(dp, profits, weights, capacity, 0)
}

func skrMemoized(dp [][]int, profits, weights []int, capacity, curIdx int) int {
	if curIdx >= len(profits) || capacity <= 0 {
		return 0
	}
	if dp[curIdx][capacity] != -1 {
		return dp[curIdx][capacity]
	}
	p1 := 0
	if weights[curIdx] <= capacity {
		p1 = profits[curIdx] + skr(profits, weights, capacity-weights[curIdx], curIdx+1)
	}
	p2 := skr(profits, weights, capacity, curIdx+1)
	dp[curIdx][capacity] = utils.Max(p1, p2)
	return dp[curIdx][capacity]
}

func solveKnapsackTabulated(profits, weights []int, capacity int) int {
	if len(profits) != len(weights) || len(profits) == 0 || capacity <= 0 {
		return 0
	}
	N := len(profits)
	dp := make([][]int, N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, capacity+1)
	}

	// 0 Capacity column of dp is already set to 0 by default

	// Initialize the array by just including first item, if it is less than capacity at that index
	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[0][c] = profits[0]
		}
	}

	for i := 1; i < N; i++ {
		for c := 1; c <= capacity; c++ {
			p1 := 0
			if weights[i] <= c {
				p1 = profits[i] + dp[i-1][c-weights[i]]
			}
			p2 := dp[i-1][c]
			dp[i][c] = utils.Max(p1, p2)
		}
	}
	// printSelectedItems(dp, profits, weights, capacity)
	return dp[N-1][capacity]
}

func printSelectedItems(dp [][]int, profits, weights []int, capacity int) {
	fmt.Print("Selected Items: ")
	N := len(profits)
	totalProfit := dp[N-1][capacity]
	for i := N - 1; i > 0; i-- {
		if totalProfit != dp[i-1][capacity] {
			fmt.Printf(" %d", weights[i])
			capacity -= weights[i]
			totalProfit -= profits[i]
		}
	}
	fmt.Println("")
}

func main() {
	p := []int{4, 5, 3, 7, 9, 2, 8, 12, 7, 1, 2, 6}
	w := []int{2, 3, 1, 4, 4, 8, 2, 7, 3, 8, 3, 3}
	t1 := time.Now()
	fmt.Println("Recursive=", SolveKnapsackRecursive(p, w, 10), time.Now().Sub(t1).Nanoseconds(), "ns")
	t1 = time.Now()
	fmt.Println("Memoized=", SolveKnapsackMemoized(p, w, 10), time.Now().Sub(t1).Nanoseconds(), "ns")
	t1 = time.Now()
	fmt.Println("Tabulated=", solveKnapsackTabulated(p, w, 10), time.Now().Sub(t1).Nanoseconds(), "ns")

}
