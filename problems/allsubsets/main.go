package main

import "fmt"

func allSubsets(arr []int) [][]int {
	result := [][]int{}
	helper([]int{}, arr, 0, &result)
	return result
}

func helper(temp, arr []int, curIdx int, result *[][]int) {
	t := []int{}
	t = append(t, temp...)
	*result = append(*result, t)
	if curIdx == len(arr) {
		return
	}
	for i := curIdx; i < len(arr); i++ {
		helper(append(temp, arr[i]), arr, i+1, result)
	}
}
func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(arr)
	fmt.Println(allSubsets(arr))
}
