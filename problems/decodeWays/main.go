package main

import (
	"fmt"
	"strconv"
)

// A message containing letters from A-Z is being encoded to numbers using the following mapping:

// 'A' -> 1
// 'B' -> 2
// ...
// 'Z' -> 26
// Given a non-empty string containing only digits, determine the total number of ways to decode it.

// Example 1:

// Input: "12"
// Output: 2
// Explanation: It could be decoded as "AB" (1 2) or "L" (12).
// Example 2:

// Input: "226"
// Output: 3
// Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).

func numDecodings(s string) int {

	if len(s) == 0 || (len(s) == 1 && s[0]-'0' != 0) {
		return len(s)
	}
	if s[0]-'0' == 0 {
		return 0
	}
	// n1 := s[0] - '0'
	n2, _ := strconv.Atoi(string(s[0:2]))
	if n2 > 26 {
		return 1 + numDecodings(s[1:])
	} else if n2 == 10 || n2 == 20 {
		return 1 + numDecodings(s[2:])
	}
	return 2 + numDecodings(s[2:])

}

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("10"))
}
