package connectedComponents

// CountComponents counts the connected components in undirected graph with n vertices and edges
func CountComponents(n int, edges [][]int) int {
	visited := make([]bool, n)
	adjList := make([][]int, n)
	for i := 0; i < len(adjList); i++ {
		adjList[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			helper(i, adjList, visited)
		}
	}
	return count
}

func helper(x int, adjList [][]int, visited []bool) {
	if visited[x] {
		return
	}
	visited[x] = true
	for i := 0; i < len(adjList[x]); i++ {
		if !visited[adjList[x][i]] {
			helper(adjList[x][i], adjList, visited)
		}
	}
}

// CountComponentsIterative counts the connected components in undirected graph with n vertices and edges iteratively
func CountComponentsIterative(n int, edges [][]int) int {
	visited := make([]bool, n)
	adjList := make([][]int, n)
	stack := []int{}
	for i := 0; i < len(adjList); i++ {
		adjList[i] = []int{}
	}
	for i := 0; i < len(edges); i++ {
		adjList[edges[i][0]] = append(adjList[edges[i][0]], edges[i][1])
		adjList[edges[i][1]] = append(adjList[edges[i][1]], edges[i][0])
	}
	count := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			count++
			stack = append([]int{i}, stack...)

			for len(stack) > 0 {
				curr := stack[0]
				stack = stack[1:]
				visited[curr] = true
				for _, neighbour := range adjList[curr] {
					if !visited[neighbour] {
						stack = append([]int{neighbour}, stack...)
					}
				}
			}
		}
	}
	return count
}
