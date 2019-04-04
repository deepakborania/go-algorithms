package main

import (
	"fmt"
	"sort"
)

func Permute(arr []int) [][]int {
	result := [][]int{}
	if len(arr) == 0 {
		return result
	}
	sort.Ints(arr)
	helper(arr, []int{}, &result, make([]bool, len(arr)))
	return result
}

func helper(arr []int, comb []int, result *[][]int, used []bool) {
	if len(comb) == len(arr) {
		*result = append(*result, comb)
		return
	}
	for i := 0; i < len(arr); i++ {
		if used[i] || i > 0 && arr[i] == arr[i-1] && !used[i-1] {
			continue
		}
		used[i] = true
		comb = append(comb, arr[i])
		helper(arr, comb, result, used)
		used[i] = false
		comb = comb[:len(comb)-1]
	}
}
func main() {
	arr := []int{1, 1, 2}
	fmt.Println(arr, "=", Permute(arr))
}
