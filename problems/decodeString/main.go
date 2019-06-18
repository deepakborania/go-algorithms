package main

import (
	"fmt"
	"strconv"
)

/*
394. Decode String
*/
func decodeString(s string) string {
	if s == "" {
		return s
	}
	prefix := ""
	repStr := ""
	start, end := 0, -1
	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			prefix += string(s[i])
		} else if s[i] >= '0' && s[i] <= '9' {
			repStr += string(s[i])
		} else if s[i] == '[' {
			start = i + 1
			count := 1
			i++
			for {
				if s[i] == '[' {
					count++
				} else if s[i] == ']' {
					count--
					if count == 0 {
						end = i - 1
						rep, _ := strconv.Atoi(repStr)
						res := prefix
						for j := 0; j < rep; j++ {
							res += decodeString(s[start : end+1])
						}
						prefix = res
						repStr = ""
						break
					}
				}
				i++
			}
		}
	}
	return prefix
}

func main() {
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("no2[abc]yoyo3[cd]ef"))
}
