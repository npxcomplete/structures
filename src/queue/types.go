package queue

type Item interface{}

type Interface interface {
	Peek() Item
	Enqueue(Item)
	Dequeue() Item
	Len() int
}
