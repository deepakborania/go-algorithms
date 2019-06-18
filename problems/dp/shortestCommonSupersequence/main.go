package main

import (
	"fmt"

	"github.com/deepakborania/go-algorithms/utils"
)

func findSCSLength(s1, s2 string) int {
	return helper(s1, s2, 0, 0)
}

func helper(s1, s2 string, i1, i2 int) int {
	if i1 == len(s1) {
		return len(s2) - i2
	} else if i2 == len(s2) {
		return len(s1) - i1
	}
	if s1[i1] == s2[i2] {
		return 1 + helper(s1, s2, i1+1, i2+1)
	}
	l1 := 1 + helper(s1, s2, i1, i2+1)
	l2 := 1 + helper(s1, s2, i1+1, i2)
	return utils.Min(l1, l2)
}

func main() {
	fmt.Println(findSCSLength("dynamic", "programming"))
}
