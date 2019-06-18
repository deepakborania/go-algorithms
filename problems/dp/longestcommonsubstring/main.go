package main

import (
	"fmt"

	"github.com/deepakborania/go-algorithms/utils"
)

func lengthLCS(s1, s2 string) int {
	m := map[string]int{}
	return helper(s1, s2, 0, 0, 0, m)
}

func helper(s1, s2 string, idx1, idx2, count int, m map[string]int) int {
	if idx1 == len(s1) || idx2 == len(s2) {
		return count
	}
	key := getKey(idx1, idx2, count)
	if v, ok := m[key]; ok {
		return v
	}
	if s1[idx1] == s2[idx2] {
		count = helper(s1, s2, idx1+1, idx2+1, count+1, m)
	}
	c1 := helper(s1, s2, idx1+1, idx2, 0, m)
	c2 := helper(s1, s2, idx1, idx2+1, 0, m)
	m[key] = utils.Max(count, utils.Max(c1, c2))
	return m[key]
}
func getKey(i1, i2, c int) string {
	return fmt.Sprintf("%d|%d|%d", i1, i2, c)
}

func lengthLCSBU(s1, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	result := 0
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				result = utils.Max(dp[i][j], result)
			}
		}
	}
	return result
}
func main() {
	fmt.Println(lengthLCS("abdca", "cbda"))      //2
	fmt.Println(lengthLCS("passport", "ppsspt")) //3

	fmt.Println(lengthLCSBU("abdca", "cbda"))      //2
	fmt.Println(lengthLCSBU("passport", "ppsspt")) //3
}
