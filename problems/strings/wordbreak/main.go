package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}
	m := map[string]bool{}
	for _, ss := range wordDict {
		m[ss] = true
	}
	return helper(s, 0, 0, m, dp)
}

func helper(s string, start, end int, m map[string]bool, dp [][]int) bool {
	if start < len(s) && end < len(s) && dp[start][end] != 0 {
		return dp[start][end] == 1
	}
	if start == len(s) {
		return true
	}
	if end >= len(s) {
		return false
	}
	if m[string(s[start:end+1])] && helper(s, end+1, end+1, m, dp) {
		dp[start][end] = 1
		return true
	}
	res := helper(s, start, end+1, m, dp)
	if res {
		dp[start][end] = 1
	} else {
		dp[start][end] = -1
	}

	return dp[start][end] == 1

}

func main() {
	s1 := "applepenapple"
	d1 := []string{"apple", "pen"}
	fmt.Println(s1, wordBreak(s1, d1))

	s2 := "catsandog"
	d2 := []string{"cats", "dog", "sand", "and", "cat"}
	fmt.Println(s2, wordBreak(s2, d2))

	s2 = "leetcode"
	d2 = []string{"leet", "code"}
	fmt.Println(s2, wordBreak(s2, d2))

	s2 = "leetscode"
	d2 = []string{"leet", "leets", "code"}
	fmt.Println(s2, wordBreak(s2, d2))
}
