package main

import "fmt"

func isBipartite(graph [][]int) bool {
	N := len(graph)
	colors := make([]int, N)
	for i := 0; i < N; i++ { //To check for disconnected graph
		if colors[i] == 0 && !validColor(graph, colors, 1, i) {
			return false
		}
	}
	return true
}

func validColor(graph [][]int, colors []int, color, node int) bool {
	if colors[node] != 0 {
		return colors[node] == color
	}
	colors[node] = color
	for _, n := range graph[node] {
		if !validColor(graph, colors, -color, n) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}
