package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type priorityqueue []int

func (pq priorityqueue) Len() int            { return len(pq) }
func (pq priorityqueue) Less(i, j int) bool  { return pq[i] < pq[j] }
func (pq priorityqueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityqueue) Push(x interface{}) { *pq = append(*pq, x.(int)) }
func (pq *priorityqueue) Pop() interface{} {
	old := *pq
	retVal := old[old.Len()-1]
	*pq = old[:old.Len()-1]
	return retVal
}

//[[0, 30],[5, 10],[15, 20]]
func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	pq := priorityqueue{}
	heap.Init(&pq)
	minRooms := 0
	for i := 0; i < len(intervals); i++ {
		if len(pq) == 0 {
			heap.Push(&pq, intervals[i][1])
		} else {
			if pq[0] <= intervals[i][0] {
				heap.Pop(&pq)
			}
			heap.Push(&pq, intervals[i][1])
		}
		minRooms = max(minRooms, len(pq))
	}
	return minRooms
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
func main() {
	fmt.Println(minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}}))
}
