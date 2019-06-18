package main

import "fmt"

func removeInvalidParentheses(s string) []string {
	q := []string{s}
	visited := map[string]bool{}
	visited[s] = true
	result := []string{}
	found := false
	for len(q) > 0 {
		e := q[0]
		q = q[1:]
		if isValid(e) {
			found = true
			result = append(result, e)
		}
		if found {
			continue
		}
		for i := 0; i < len(e); i++ {
			if e[i] != '(' && e[i] != ')' {
				continue
			}
			str := e[:i] + e[i+1:]
			if !visited[str] {
				q = append(q, str)
				visited[str] = true
			}

		}
	}
	return result
}

func isValid(s string) bool {
	stack := 0
	for _, c := range s {
		if c == '(' {
			stack++
		} else if c == ')' {
			stack--
		}
		if stack < 0 {
			return false
		}
	}
	return stack == 0
}

func main() {
	fmt.Println(removeInvalidParentheses(")("))
}
