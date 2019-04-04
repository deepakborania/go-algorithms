package main

import (
	"fmt"

	"github.com/deepakborania/go-algorithms/utils"
)

// MaxSubarraySum return the sum of maximum subarray
func MaxSubarraySum(arr []int) int {
	sumHere, max := 0, 0
	for i := 0; i < len(arr); i++ {
		// Is the current element larger than adding current element to sum?
		// Yes: Restart the sum here (When previous sum is -ve and this elem is +ve)
		//No: Add this element to sumHere
		sumHere = utils.Max(arr[i], sumHere+arr[i])

		// Update max with max sumHere
		max = utils.Max(sumHere, max)
	}
	return max

}

func main() {
	arr := []int{-1, 2, 4, -3, 5, 2, -5, 2}
	fmt.Println(arr, "=", MaxSubarraySum(arr))
}
