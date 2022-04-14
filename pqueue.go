package kyu

import "container/heap"

// PQueue is a priority queue.
//
// It supports efficient inspection and removal of elements at the front of the
// queue. The lowest elements appear towards the front of the queue.
type PQueue struct {
	// Less returns true if a should be closer to the front of the queue than b.
	Less func(a, b any) bool

	h pheap
}

// Len returns the number of elements in the queue.
func (q *PQueue) Len() int {
	return q.h.Len()
}

// Push adds a new value to the queue.
//
// It returns the element that contains that value.
func (q *PQueue) Push(v any) *Element {
	if q.h.elements == nil {
		q.h.less = q.Less
	}

	e := &Element{Value: v}
	heap.Push(&q.h, e)

	return e
}

// Peek returns the element at the front of the queue without removing it from
// the queue.
//
// If the queue is empty, e is nil and ok is false.
func (q *PQueue) Peek() (e *Element, ok bool) {
	if q.h.Len() == 0 {
		return nil, false
	}

	return q.h.elements[0], true
}

// Pop removes the element at the front of the queue and returns its value.
//
// If the queue is empty, v is nil and ok is false.
func (q *PQueue) Pop() (v any, ok bool) {
	if q.h.Len() == 0 {
		return nil, false
	}

	v = q.h.elements[0].Value
	heap.Pop(&q.h)

	return v, true
}

// Contains returns true if e is in the queue.
func (q *PQueue) Contains(e *Element) bool {
	if e.index < 0 || e.index >= q.h.Len() {
		return false
	}

	return q.h.elements[e.index] == e
}

// IsFront returns true if e is at the front of the queue.
func (q *PQueue) IsFront(e *Element) bool {
	if q.h.Len() == 0 {
		return false
	}

	return q.h.elements[0] == e
}

// Update reorders the queue to reflect a change in e.Value that might cause e
// to occupy a different position within in the queue.
func (q *PQueue) Update(e *Element) {
	if !q.Contains(e) {
		panic("element is not on the queue")
	}

	heap.Fix(&q.h, e.index)
}

// Remove removes e from the queue.
func (q *PQueue) Remove(e *Element) {
	if q.Contains(e) {
		heap.Remove(&q.h, e.index)
	}
}
