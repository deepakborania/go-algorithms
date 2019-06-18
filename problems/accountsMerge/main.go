package main

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	//map[string]set
	graph := map[string]map[string]bool{}
	for _, ls := range accounts {
		for i := 1; i < len(ls); i++ {
			if _, ok := graph[ls[i]]; !ok {
				graph[ls[i]] = map[string]bool{}
			}
			graph[ls[i]][ls[1]] = true
			graph[ls[1]][ls[i]] = true
		}
	}
	visited := map[string]bool{}
	result := [][]string{}
	for _, ls := range accounts {
		if !visited[ls[1]] {
			ans := []string{}
			dfs(graph, visited, ls[1], &ans)
			sort.Strings(ans)
			ans = append([]string{ls[0]}, ans...)
			result = append(result, ans)
		}
	}
	return result
}

func dfs(graph map[string]map[string]bool, visited map[string]bool, s string, ans *[]string) {
	*ans = append(*ans, s)
	visited[s] = true
	for _, str := range keys(graph[s]) {
		if !visited[str] {
			dfs(graph, visited, str, ans)
		}
	}
}
func keys(g map[string]bool) []string {
	keys := []string{}
	for k := range g {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	fmt.Println(accountsMerge([][]string{
		{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"},
		{"John", "johnsmith@mail.com", "john_newyork@mail.com"},
		{"Mary", "mary@mail.com"},
	}))
}
