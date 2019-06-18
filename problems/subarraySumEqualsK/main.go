package main

import "fmt"

func subarraySum(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				count++
			}
		}
	}
	return count
}

func main() {
	fmt.Println(subarraySum([]int{-1, 0, 1, -1, 0, 1}, 0))
}
