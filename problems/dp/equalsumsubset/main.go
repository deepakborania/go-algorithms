package main

import "fmt"

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	// We are adding two equal halves so sum can only be even
	if sum%2 != 0 {
		return false
	}
	//0=nil, 1=false, 2=true
	dp := make([][]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, (sum/2)+1)
	}
	return helper(dp, nums, sum/2, 0)
}

func helper(dp [][]int, nums []int, sum int, curIdx int) bool {
	if sum == 0 {
		return true
	}
	if curIdx >= len(nums) || len(nums) == 0 {
		return false
	}
	if dp[curIdx][sum] != 0 {
		return dp[curIdx][sum] == 2
	}
	if nums[curIdx] <= sum {
		if helper(dp, nums, sum-nums[curIdx], curIdx+1) {
			dp[curIdx][sum] = 2
			return true
		}
	}
	result := helper(dp, nums, sum, curIdx+1)
	if result == true {
		dp[curIdx][sum] = 2
	} else {
		dp[curIdx][sum] = 1
	}
	return dp[curIdx][sum] == 2
}

func canPartitionBU(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	N := len(nums)
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]bool, sum+1)
	}
	// Set first row of table. If nums[0] <= sum@s, set true
	for s := 0; s <= sum; s++ {
		dp[0][s] = nums[0] == s
	}

	for i := 1; i < N; i++ {
		for s := 1; s <= sum; s++ {
			// If we can get sum s without num at index i
			if dp[i-1][s] {
				dp[i][s] = dp[i-1][s]
			} else if s >= nums[i] {
				dp[i][s] = dp[i-1][s-nums[i]]
			}
			// If solution achieved faster
			if s == sum && dp[i][s] {
				return true
			}
		}
	}
	return dp[N-1][sum]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))         //true
	fmt.Println(canPartition([]int{1, 2, 3, 5}))          //false
	fmt.Println(canPartition([]int{1, 2, 3, 4, 5, 6, 7})) //true
	//false
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100}
	fmt.Println(canPartition(arr))
	fmt.Println(canPartitionBU(arr))
}
