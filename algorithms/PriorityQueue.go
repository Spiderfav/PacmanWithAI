package algorithms

import (
	"container/heap"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

// This priority queue is based on the Golang's own documentation from the GO heap package:
// https://pkg.go.dev/container/heap .
// However, to fit my purposes, as I will use the distance as priority, some functions have changed.
type PriorityNode struct {
	node     mazegrid.Position
	priority float64
	index    int
}

// A PriorityQueue implements heap.Interface and holds PriorityNodes
type PriorityQueue []*PriorityNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority/distance so we use less than here
	// Low distance to destination = higher priority
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityNode)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// Update changes the priority and value of an PriorityNode in the queue
func (pq *PriorityQueue) update(item *PriorityNode, value mazegrid.Position, priority float64) {
	item.node = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
