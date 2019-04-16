package main

import "fmt"

func permuteString(s string) []string {
	accum := []string{}
	permute("", s, &accum)
	return accum
}
func permute(prefix, suffix string, accum *[]string) {
	N := len(suffix)
	if N == 0 {
		*accum = append(*accum, prefix)
		return
	}
	for i := 0; i < N; i++ {
		permute(prefix+string(suffix[i]), suffix[:i]+suffix[i+1:], accum)
	}
}
func main() {
	fmt.Println(permuteString("abcd"))
}
