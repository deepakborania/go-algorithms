package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	heights := [][]int{}
	results := [][]int{}
	for _, b := range buildings {
		heights = append(heights, []int{b[0], -b[2]})
		heights = append(heights, []int{b[1], b[2]})
	}

	sort.Slice(heights, func(i, j int) bool {
		if heights[i][0] != heights[j][0] {
			return heights[i][0] < heights[j][0]
		} else {
			return heights[i][1] < heights[j][1]
		}
	})
	pq := &priorityQueue{0}
	heap.Init(pq)
	prev := 0
	for _, h := range heights {
		if h[1] < 0 {
			heap.Push(pq, -h[1])
		} else {
			for i := 0; i < len(*pq); i++ {
				if (*pq)[i] == h[1] {
					heap.Remove(pq, i)
					break
				}
			}
		}
		cur := heap.Pop(pq).(int) //emulate peek
		heap.Push(pq, cur)
		if cur != prev {
			results = append(results, []int{h[0], cur})
			prev = cur
		}
	}
	return results
}

type priorityQueue []int

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i] > pq[j] }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	fmt.Println(getSkyline([][]int{{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8}}))
}
