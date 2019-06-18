package main

import "fmt"

func subsets(arr []int) [][]int {
	result := [][]int{}
	helper(arr, 0, []int{}, &result)
	return result
}
func helper(nums []int, idx int, curr []int, result *[][]int) {
	*result = append(*result, append([]int{}, curr...))
	for i := idx; i < len(nums); i++ {
		curr = append(curr, nums[i])
		helper(nums, i+1, curr, result)
		curr = curr[:len(curr)-1]
	}
}
func main() {
	// fmt.Println(subsets([]int{3, 0, 1, 2}))
	
}
