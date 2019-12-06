package templates

import (
	"github.com/cheekybits/genny/generic"

	"github.com/npxcomplete/structures/src/queue"
)

type GenericItem generic.Type

func NewGenericItemPriorityQueue(capacity int, niceness func(GenericItem) int) privateGenericItemPriorityQueue {
	return privateGenericItemPriorityQueue{queue.NewPriorityQueue(capacity, func(item queue.Item) int { niceness(item.(GenericItem)) })}
}

// genny is case sensitive even though this has other meanings in go, so we prefix the intent.
type privateGenericItemPriorityQueue struct {
	generic queue.Interface
}

func (pq privateGenericItemPriorityQueue) Peek() GenericItem {
	result, _ := pq.generic.Peek().(GenericItem)
	return result
}

func (pq privateGenericItemPriorityQueue) Enqueue(value GenericItem) {
	pq.generic.Enqueue(value)
}

func (pq privateGenericItemPriorityQueue) Dequeue() GenericItem {
	result, _ := pq.generic.Dequeue().(GenericItem)
	return result
}

func (pq privateGenericItemPriorityQueue) Len() int {
	return pq.generic.Len()
}
