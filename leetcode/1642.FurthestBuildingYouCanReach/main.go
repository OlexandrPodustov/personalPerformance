package main

import (
	"container/heap"
)

// https://pkg.go.dev/container/heap#example-package-IntHeap
// // This example inserts several ints into an IntHeap, checks the minimum,
// // and removes them in order of priority.
// func main() {
// 	h := &IntHeap{2, 1, 5}
// 	heap.Init(h)
// 	heap.Push(h, 3)
// 	fmt.Printf("minimum: %d\n", (*h)[0])
// 	for h.Len() > 0 {
// 		fmt.Printf("%d ", heap.Pop(h))
// 	}
// }

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	queue := &IntHeap{}
	heap.Init(queue)

	for i := range heights[:len(heights)-1] {
		diff := heights[i+1] - heights[i]
		if diff < 0 {
			continue
		}
		heap.Push(queue, diff)
		if queue.Len() > ladders {
			lowestAvailableJump := heap.Pop(queue)
			laj, _ := lowestAvailableJump.(int)
			bricks -= laj

			if bricks < 0 {
				return i
			}
		}
	}

	return len(heights) - 1
}
