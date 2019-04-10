package utils

import "fmt"

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func PrintMatrixBool(arr [][]bool) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%5v ", arr[i][j])
		}
		fmt.Println("")
	}
}

func PrintMatrixInt(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%3d ", arr[i][j])
		}
		fmt.Println("")
	}
}
