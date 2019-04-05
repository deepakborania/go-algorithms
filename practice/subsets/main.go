package main

import "fmt"

func Subsets(arr []int) [][]int {
	result := [][]int{}
	if len(arr) == 0 {
		return result
	}
	ss(arr, []int{}, 0, &result)
	return result
}

func ss(arr []int, comb []int, idx int, result *[][]int) {
	*result = append(*result, append([]int{}, comb...))
	for i := idx; i < len(arr); i++ {
		comb = append(comb, arr[i])
		ss(arr, comb, i+1, result)
		comb = comb[:len(comb)-1]
	}
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr, "=", Subsets(arr))
}
