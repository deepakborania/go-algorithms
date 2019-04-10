package main

import "fmt"

// We are given a set of +ve numbers and we need to determine if there exists a subset whose sum equal to given s.
func subsetSum(nums []int, s int) bool {
	if s > 0 && len(nums) == 0 {
		return false
	}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum < s {
		return false
	} else if sum == s {
		return true
	}
	return helper(nums, s, 0)
}

func helper(nums []int, sum, curIdx int) bool {
	if sum == 0 {
		return true
	}
	if curIdx >= len(nums) {
		return false
	}
	if nums[curIdx] <= sum {
		if helper(nums, sum-nums[curIdx], curIdx+1) {
			return true
		}
	}
	return helper(nums, sum, curIdx+1)
}

func subsetSumBU(nums []int, S int) bool {
	if S > 0 && len(nums) == 0 {
		return false
	}
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum < S {
		return false
	} else if sum == S {
		return true
	}
	dp := make([]bool, S+1)
	dp[0] = true
	for i := 0; i < len(nums); i++ {
		for s := S; s >= nums[i]; s-- {
			dp[s] = dp[s] || dp[s-nums[i]]
			if s == S && dp[s] {
				return true
			}
		}
	}
	return dp[S]
}
func main() {
	fmt.Println(subsetSum([]int{1, 2, 3, 7}, 6))
	fmt.Println(subsetSum([]int{1, 2, 7, 1, 5}, 10))
	fmt.Println(subsetSum([]int{1, 3, 4, 8}, 6))

	fmt.Println(subsetSumBU([]int{1, 2, 3, 7}, 6))
	fmt.Println(subsetSumBU([]int{1, 2, 7, 1, 5}, 10))
	fmt.Println(subsetSumBU([]int{1, 3, 4, 8}, 6))
}
