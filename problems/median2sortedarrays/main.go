package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var X, Y []int

	if len(nums1) <= len(nums2) {
		X, Y = nums1, nums2
	} else {
		Y, X = nums1, nums2
	}
	start, end := 0, len(X)

	for start <= end {
		partitionX := (start + end) / 2
		partitionY := ((len(X) + len(Y) + 1) / 2) - partitionX

		leftMaxX, rightMinX := math.MinInt64, math.MaxInt64
		leftMaxY, rightMinY := math.MinInt64, math.MaxInt64
		if partitionX > 0 {
			leftMaxX = X[partitionX-1]
		}
		if partitionX < len(X) {
			rightMinX = X[partitionX]
		}
		if partitionY > 0 {
			leftMaxY = Y[partitionY-1]
		}
		if partitionY < len(Y) {
			rightMinY = Y[partitionY]
		}
		fmt.Println(leftMaxX, rightMinX, leftMaxY, rightMinY)
		if leftMaxX <= rightMinY && leftMaxY <= rightMinX {
			if (len(X)+len(Y))%2 != 0 {
				return float64(max(leftMaxX, leftMaxY))
			} else {
				return float64(max(leftMaxX, leftMaxY)+min(rightMinX, rightMinY)) / 2
			}
		} else if leftMaxX > rightMinY {
			end = partitionX - 1
		} else {
			start = partitionX + 1
		}
	}
	return 0.0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}
