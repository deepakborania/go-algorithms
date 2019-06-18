package main

import "fmt"

func countOfSubsetSums(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum == S {
		return 1
	}
	if sum < S {
		return 0
	}
	dp := make([][]*int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]*int, S+1)
	}
	return helper(nums, S, 0, dp)
}

func helper(nums []int, s, curIdx int, dp [][]*int) int {
	if s == 0 {
		return 1
	}
	if curIdx >= len(nums) {
		return 0
	}
	if dp[curIdx][s] != nil {
		return *dp[curIdx][s]
	}
	sum1 := 0
	if nums[curIdx] <= s {
		sum1 = helper(nums, s-nums[curIdx], curIdx+1, dp)
	}
	sum2 := helper(nums, s, curIdx+1, dp)
	res := sum1 + sum2
	dp[curIdx][s] = &res
	return res
}
func main() {
	arr := []int{1, 1, 2, 3}
	fmt.Println(arr, countOfSubsetSums(arr, 4)) //ans = 3

	arr = []int{1, 2, 7, 1, 5}
	fmt.Println(arr, countOfSubsetSums(arr, 9)) //ans = 3

	// Will be terribily slow on non-memoized recursion
	arr = []int{1, 7, 6, 5, 4, 3, 2, 1, 2, 3, 2, 7, 1, 5, 7, 6, 3, 8, 4, 9, 3, 5, 7, 3, 7, 8, 3, 0, 1, 9, 2, 8, 3, 4, 7, 8, 7, 3, 8, 2, 8, 2, 7, 3, 7, 2, 4, 5, 7, 7, 9, 8, 5, 2, 3, 6, 7}
	fmt.Println(arr, countOfSubsetSums(arr, 45)) //ans = 42044074666
}
