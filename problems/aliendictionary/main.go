package main

import "fmt"

func alienOrder(words []string) string {
	graph := map[byte]map[byte]bool{}
	inDegree := make([]int, 26)

	buildGraph(words, graph, inDegree)
	order := topologicalSort(graph, inDegree)
	return order
}

func buildGraph(words []string, graph map[byte]map[byte]bool, inDegree []int) {
	// Create an empty graph with just the chars initialized
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			graph[words[i][j]] = map[byte]bool{}
		}
	}
	//Construct graph and indegree array
	for i := 1; i < len(words); i++ {
		first := words[i-1]
		second := words[i]
		length := min(len(first), len(second))

		for j := 0; j < length; j++ {
			parent, child := first[j], second[j]
			if first[j] != second[j] {
				if !graph[parent][child] {
					graph[parent][child] = true
					inDegree[child-'a']++
				}
				break
			}
		}
	}
}

func topologicalSort(graph map[byte]map[byte]bool, inDegree []int) string {
	q := []byte{}
	for _, _b := range getAllKeys(graph) {
		if inDegree[_b-'a'] == 0 {
			q = append(q, _b)
		}
	}
	result := ""
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		result += string(v)
		for neighbour, _ := range graph[v] {
			inDegree[neighbour-'a']--
			if inDegree[neighbour-'a'] == 0 {
				q = append(q, neighbour)
			}
		}
	}
	if len(graph) != len(result) {
		return ""
	}
	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getAllKeys(graph map[byte]map[byte]bool) []byte {
	keys := []byte{}
	for k := range graph {
		keys = append(keys, k)
	}
	return keys
}
func main() {
	fmt.Println(alienOrder([]string{"wrt", "wrf", "er", "ett", "rftt"}))
	fmt.Println(alienOrder([]string{"ri", "xz", "qxf", "jhsguaw", "dztqrbwbm", "dhdqfb", "jdv", "fcgfsilnb", "ooby"}))
}
