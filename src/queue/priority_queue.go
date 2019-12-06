package queue

import (
	"container/heap"
	"math"
)

var _ Interface = &priorityQueue{}

func NewPriorityQueue(capacity int, niceness func(Item) int) *priorityQueue {
	return &priorityQueue{
		niceness:  niceness,
		heapSlice: make([]*item, 0, int(math.Max(1, float64(capacity)))),
		autoscale: capacity > 0,
	}
}

type priorityQueue struct {
	// Similar to the nice function in unix low values
	// are high priority, high values are nicer and let
	// others cut in line.
	niceness  func(Item) int
	heapSlice heapSlice
	autoscale bool
}

type item struct {
	value    Item // The value of the item; arbitrary.
	priority int  // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

func (pq *priorityQueue) Peek() Item {
	return pq.heapSlice[0].value
}

func (pq *priorityQueue) Enqueue(val Item) {
	n := len(pq.heapSlice)

	pq.heapSlice = append(pq.heapSlice, &item{
		value:    val,
		priority: pq.niceness(val),
		index:    n,
	})
	heap.Fix(&pq.heapSlice, n)
}

func (pq *priorityQueue) Dequeue() Item {
	val, _ := heap.Pop(&pq.heapSlice).(*item)
	return val.value
}

func (pq *priorityQueue) Len() int {
	return pq.heapSlice.Len()
}

////////////// Type over which the heap algorithm performs /////////

type heapSlice []*item

var heapSliceTypeAssertion heapSlice = []*item{}
var _ heap.Interface = &heapSliceTypeAssertion

func (hs *heapSlice) Push(x interface{}) {
	n := len(*hs)
	item := x.(*item)
	item.index = n
	*hs = append(*hs, item)
}

func (hs *heapSlice) Pop() interface{} {
	old := *hs
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*hs = old[0 : n-1]
	return item
}

func (hs *heapSlice) update(item *item) {
	heap.Fix(hs, item.index)
}

////////////// sort.Interface //////////////////////////

func (hs heapSlice) Len() int { return len(hs) }

func (hs heapSlice) Less(i, j int) bool {
	return hs[i].priority < hs[j].priority
}

func (hs heapSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
	hs[i].index = i
	hs[j].index = j
}
