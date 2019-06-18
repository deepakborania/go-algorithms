package main

import "fmt"

func generateParenthesis(n int) []string {
	result := []string{}
	helper("", 0, 0, n, &result)
	return result
}

func helper(curr string, open, close, n int, result *[]string) {
	if len(curr) == n*2 {
		*result = append(*result, curr)
	} else {
		if open < n {
			helper(curr+"(", open+1, close, n, result)
		}
		if close < open {
			helper(curr+")", open, close+1, n, result)
		}
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
}
