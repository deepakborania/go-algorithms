package main

import (
	"fmt"
	"sort"
)

// Leet Code 698
// Given an array of integers nums and a positive integer k,
// find whether it's possible to divide this array into k non-empty subsets whose sums are all equal.
func canPartitionKSubsets(nums []int, k int) bool {
	totalSum := 0
	// Find sum of elements
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	// If totalSum/k is not integer, return false
	if totalSum%k > 0 {
		return false
	}
	// target is the sum we need
	target := totalSum / k
	sort.Ints(nums)
	idx := len(nums) - 1
	// if any element is larger than target it means target cannot be completed
	if nums[idx] > target {
		return false
	}
	// reduce number of required subsets by the number of elems which are exactly eual to target
	for i := idx; i >= 0 && nums[i] == target; i-- {
		k--
		idx--
	}
	// groups has the subset sum
	groups := make([]int, k)
	return helper(nums, target, idx, groups)
}

func helper(nums []int, target, idx int, groups []int) bool {
	if idx < 0 {
		return true
	}
	// element to consider
	elem := nums[idx]
	idx--
	for i := 0; i < len(groups); i++ {
		if groups[i]+elem <= target {
			// Try adding elem to this group
			groups[i] += elem
			// Recurse on shorter nums. Reduced previously by idx--
			if helper(nums, target, idx, groups) {
				return true
			}
			// In case last selection of group wasn't right for this element, backtrack
			groups[i] -= elem
		}
		//Ensure all 0 groups are at the end. Reduces recursive calls.
		// Infact can break for any repeated values of groups[i]
		if groups[i] == 0 {
			break
		}
	}
	return false
}
func main() {
	nums := []int{4, 3, 2, 3, 5, 2, 1}
	fmt.Println(nums, "=", canPartitionKSubsets(nums, 4))

	nums = []int{1, 5, 11, 5}
	fmt.Println(nums, "=", canPartitionKSubsets(nums, 2))

	nums = []int{1, 2, 3, 5}
	fmt.Println(nums, "=", canPartitionKSubsets(nums, 2))
}
