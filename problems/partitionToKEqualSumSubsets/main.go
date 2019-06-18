package main

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%k != 0 {
		return false
	}
	targetSum := sum / k
	sort.Ints(nums)
	if nums[len(nums)-1] > targetSum {
		return false
	}
	idx := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == targetSum {
			k--
			idx--
		}
	}
	if k == 0 {
		return true
	}
	return helper(nums[:idx], k, 0, targetSum, make([]bool, idx), 0)

}

func helper(nums []int, k, currSum, targetSum int, visited []bool, pos int) bool {
	if k == 0 {
		return true
	} else if currSum > targetSum {
		return false
	} else if currSum == targetSum {
		return helper(nums, k-1, 0, targetSum, visited, 0)
	}

	for i := pos; i < len(nums); i++ {
		if !visited[i] {
			visited[i] = true
			if helper(nums, k, currSum+nums[i], targetSum, visited, i+1) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}
func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
