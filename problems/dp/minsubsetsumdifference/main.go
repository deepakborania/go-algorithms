package main

import (
	"fmt"

	"github.com/deepakborania/go-algorithms/utils"
)

// Given a set of positive numbers, partition the set into two subsets with a minimum difference between their subset sums.
// Input: {1, 2, 3, 9}
// Output: 3
// Explanation: We can partition the given set into two subsets where minimum absolute difference
// between the sum of numbers is '3'. Following are the two subsets: {1, 2, 3} & {9}.
func minSubsetSumDiff(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, sum+1)
	}
	return helper(dp, nums, 0, 0, 0)
}
func helper(dp [][]int, nums []int, sum1, sum2, curIdx int) int {
	if curIdx >= len(nums) {
		return utils.Abs(sum1 - sum2)
	}
	if dp[curIdx][sum1] > 0 {
		return dp[curIdx][sum1]
	}
	diff1 := helper(dp, nums, sum1+nums[curIdx], sum2, curIdx+1)
	diff2 := helper(dp, nums, sum1, sum2+nums[curIdx], curIdx+1)
	dp[curIdx][sum1] = utils.Min(diff1, diff2)
	return dp[curIdx][sum1]
}

func minSubsetSumDiffBU(nums []int) int {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	N := len(nums)
	dp := make([][]bool, N)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, sum/2+1)
	}
	for i := 0; i < N; i++ {
		dp[i][0] = true
	}
	for s := 0; s <= sum/2; s++ {
		if nums[0] == s {
			dp[0][s] = true
		}
	}
	for i := 1; i < N; i++ {
		for s := 1; s <= sum/2; s++ {
			if dp[i-1][s] {
				dp[i][s] = true
			} else if nums[i] <= s {
				dp[i][s] = dp[i-1][s-nums[i]]
			}
		}
		// utils.PrintMatrixBool(dp)
		// fmt.Println("")
	}
	sum1 := 0
	for i := sum / 2; i >= 0; i-- {
		if dp[N-1][i] {
			sum1 = i
			break
		}
	}
	return sum - (2 * sum1)
}

func main() {

	arr := []int{1, 2, 3, 9} //ans=3
	fmt.Println(arr, minSubsetSumDiff(arr))
	fmt.Println(arr, minSubsetSumDiffBU(arr))
	arr = []int{1, 2, 7, 1, 5} //ans=0
	fmt.Println(arr, minSubsetSumDiff(arr))
	fmt.Println(arr, minSubsetSumDiffBU(arr))
	arr = []int{1, 3, 100, 4} // ans =92
	fmt.Println(arr, minSubsetSumDiff(arr))
	fmt.Println(arr, minSubsetSumDiffBU(arr))
	arr = []int{1, 3, 100, 4, 2, 6, 32, 8, 3, 8, 3, 8, 3, 4, 7, 3, 97, 34, 4, 62, 7, 3, 9, 5, 8, 5, 78, 3, 7, 3, 7, 3, 6, 9, 5, 3, 76, 9, 5}
	fmt.Println(arr, minSubsetSumDiff(arr))
	fmt.Println(arr, minSubsetSumDiffBU(arr))
}
