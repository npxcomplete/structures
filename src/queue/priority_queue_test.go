package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_queue_ordering(t *testing.T) {
	q := NewPriorityQueue(0, func(i Item) int { return i.(int) })
	q.Enqueue(5)
	q.Enqueue(9)
	q.Enqueue(2)

	assert.Equal(t, 2, q.Dequeue())
	assert.Equal(t, 5, q.Dequeue())
	assert.Equal(t, 9, q.Dequeue())
}

func Test_peek_pop_consistent(t *testing.T) {
	q := NewPriorityQueue(0, func(i Item) int { return i.(int) })
	q.Enqueue(5)
	q.Enqueue(9)
	q.Enqueue(2)

	var lookahead Item

	lookahead = q.Peek()
	assert.Equal(t, q.Dequeue(), lookahead)
	lookahead = q.Peek()
	assert.Equal(t, q.Dequeue(), lookahead)
	lookahead = q.Peek()
	assert.Equal(t, q.Dequeue(), lookahead)
}
