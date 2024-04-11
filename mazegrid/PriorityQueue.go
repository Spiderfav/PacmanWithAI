package mazegrid

import (
	"container/heap"
)

// This priority queue is based on the Golang's own documentation from the GO heap package:
// https://pkg.go.dev/container/heap .
// However, to fit my purposes, as I will use the distance as priority, some functions have changed.
type PriorityNode struct {
	node     *MazeSquare
	priority float64
	index    int
}

// For testing purposes, added an initialisation function to the node
func (pn *PriorityNode) Init(node *MazeSquare, priority float64) {
	pn.node = node
	pn.priority = priority

}

// Returns the mazeSquare of a Node in priority queue
func (pn *PriorityNode) GetNode() *MazeSquare {
	return pn.node
}

// Returns the index of a Node in priority queue
func (pn *PriorityNode) GetIndex() int {
	return pn.index
}

// Returns the priority of a Node in priority queue
func (pn *PriorityNode) GetPriority() float64 {
	return pn.priority
}

// A PriorityQueue implements heap.Interface and holds PriorityNodes
type PriorityQueue []*PriorityNode

// Returns the length of a priority queue
func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority/distance so we use less than here
	// Low distance to destination = higher priority
	return pq[i].priority < pq[j].priority
}

// Swaps two items in the priority queue, given their indexes
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Adds a new item to the priority queue
func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*PriorityNode)
	item.index = n
	*pq = append(*pq, item)
}

// Removes the top item from the priority queue
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
func (pq *PriorityQueue) update(item *PriorityNode, value *MazeSquare, priority float64) {
	item.node = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
