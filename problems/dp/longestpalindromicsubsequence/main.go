package main

import (
	"fmt"

	"github.com/deepakborania/go-algorithms/utils"
)

func findLPSLength(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s))
	}
	return helper(s, 0, len(s)-1, dp)
}
func helper(s string, startIdx, endIdx int, dp [][]int) int {
	if startIdx > endIdx {
		return 0
	}
	// Every sequence of 1 length is a palindrome
	if startIdx == endIdx {
		return 1
	}
	if dp[startIdx][endIdx] != 0 {
		return dp[startIdx][endIdx]
	}
	if s[startIdx] == s[endIdx] {
		res := 2 + helper(s, startIdx+1, endIdx-1, dp)
		dp[startIdx][endIdx] = res
		return res
	}
	c1 := helper(s, startIdx+1, endIdx, dp)
	c2 := helper(s, startIdx, endIdx-1, dp)
	res := utils.Max(c1, c2)
	dp[startIdx][endIdx] = res
	return res
}

func findLPSLengthBU(s string) int {

}
func main() {
	s := "abdbca"
	fmt.Println(s, findLPSLength(s))
	s = "cddpd"
	fmt.Println(s, findLPSLength(s))
	s = "pqr"
	fmt.Println(s, findLPSLength(s))
	s = "euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew"
	fmt.Println(s, findLPSLength(s))
}
