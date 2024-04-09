package testing

import (
	"container/heap"
	"fmt"
	"testing"

	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/algorithms"
	"gitlab.cim.rhul.ac.uk/zkac432/PROJECT/mazegrid"
)

func TestPriorityQueue_Len(t *testing.T) {
	pq := make(algorithms.PriorityQueue, 3)
	if pq.Len() != 3 {
		t.Errorf("Expected length 3, got %d", pq.Len())
	}
}

func TestPriorityQueue_Less(t *testing.T) {
	blankSquare := mazegrid.CreateBlankSquare(20)

	node1 := &algorithms.PriorityNode{}
	node1.Init(&blankSquare, 1.0)

	node2 := &algorithms.PriorityNode{}
	node2.Init(&blankSquare, 2.5)

	priorityQueue := make(algorithms.PriorityQueue, 0)
	heap.Init(&priorityQueue)

	heap.Push(&priorityQueue, node1)

	heap.Push(&priorityQueue, node2)

	if !priorityQueue.Less(0, 1) {
		t.Errorf("Expected true for pq.Less(0, 1)")
	}
}

func TestPriorityQueue_Swap(t *testing.T) {

	blankSquare := mazegrid.CreateBlankSquare(20)

	node1 := &algorithms.PriorityNode{}
	node1.Init(&blankSquare, 1)

	node2 := &algorithms.PriorityNode{}
	node2.Init(&blankSquare, 2)

	pq := algorithms.PriorityQueue{
		node1,
		node2,
	}

	pq.Swap(0, 1)

	if pq[0].GetIndex() != 0 || pq[1].GetIndex() != 1 {
		t.Errorf("Swap did not update indices correctly")
	}
}

func TestPriorityQueue_Push(t *testing.T) {

	blankSquare := mazegrid.CreateBlankSquare(20)

	node1 := &algorithms.PriorityNode{}
	node1.Init(&blankSquare, 5)

	pq := make(algorithms.PriorityQueue, 0)

	pq.Push(node1)

	if pq.Len() != 1 {
		t.Errorf("Expected queue length of 1, got %d", pq.Len())
	}
}

func TestPriorityQueue_Pop(t *testing.T) {
	blankSquare := mazegrid.CreateBlankSquare(20)

	node1 := &algorithms.PriorityNode{}
	node1.Init(&blankSquare, 2)

	node2 := &algorithms.PriorityNode{}
	node2.Init(&blankSquare, 1)

	pq := algorithms.PriorityQueue{
		node1,
		node2,
	}

	item := pq.Pop().(*algorithms.PriorityNode)

	fmt.Println("This is the item: ", item)

	if item.GetPriority() != 1 {
		t.Errorf("Pop did not return the item with the lowest priority")
	}

	if pq.Len() != 1 {
		t.Errorf("Pop did not remove the item from the queue")
	}
}
