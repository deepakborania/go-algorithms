package main

import "fmt"

func topologicalSort(N int, adj [][]bool) []int {
	T := []int{}
	visited := make([]bool, N)
	inDegree := make([]int, N)
	q := []int{}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if adj[i][j] {
				inDegree[j]++
			}
		}
	}
	for i := 0; i < N; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
			visited[i] = true
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		T = append(T, v)
		for j := 0; j < N; j++ {
			if adj[v][j] && !visited[j] {
				inDegree[j]--
				if inDegree[j] == 0 {
					q = append(q, j)
					visited[j] = true
				}
			}
		}
	}
	return T
}

func main() {
	adj := [][]bool{
		{false, true, true, false},
		{false, false, true, true},
		{false, false, false, true},
		{false, false, false, false},
	}
	fmt.Println(topologicalSort(4, adj))

}
