package kyu

import (
	"github.com/dogmatiq/kyu/mmheap"
)

// PDeque is a double-ended priority queue.
//
// It supports efficient inspection and removal of elements at both the front
// and back of the queue. The lowest elements appear towards the front of the
// queue.
type PDeque struct {
	// Less returns true if a should be closer to the front of the queue than b.
	Less func(a, b interface{}) bool

	h pheap
}

// Len returns the number of elements in the queue.
func (q *PDeque) Len() int {
	return q.h.Len()
}

// Push adds a new value to the queue.
//
// It returns the element that contains that value.
func (q *PDeque) Push(v interface{}) *Element {
	if q.h.elements == nil {
		q.h.less = q.Less
	}

	e := &Element{Value: v}
	mmheap.Push(&q.h, e)

	return e
}

// Peek returns the element at the front of the queue without removing it from
// the queue.
//
// If the queue is empty, e is nil and ok is false.
func (q *PDeque) Peek() (e *Element, ok bool) {
	if q.h.Len() == 0 {
		return nil, false
	}

	return q.h.elements[0], true
}

// PeekBack returns the element at the back of the queue without removing it
// from the queue.
//
// If the queue is empty, e is nil and ok is false.
func (q *PDeque) PeekBack() (e *Element, ok bool) {
	i := mmheap.Max(&q.h)
	if i == -1 {
		return nil, false
	}

	return q.h.elements[i], true
}

// Pop removes the element at the front of the queue and returns its value.
//
// If the queue is empty, v is nil and ok is false.
func (q *PDeque) Pop() (v interface{}, ok bool) {
	if q.h.Len() == 0 {
		return nil, false
	}

	v = q.h.elements[0].Value
	mmheap.Pop(&q.h)

	return v, true
}

// PopBack removes the element at the back of the queue and returns its value.
//
// If the queue is empty, v is nil and ok is false.
func (q *PDeque) PopBack() (v interface{}, ok bool) {
	i := mmheap.Max(&q.h)
	if i == -1 {
		return nil, false
	}

	v = q.h.elements[i].Value
	mmheap.PopMax(&q.h)

	return v, true
}

// IsFront returns true if e is at the front of the queue.
func (q *PDeque) IsFront(e *Element) bool {
	return q.h.elements[0] == e
}

// IsBack returns true if e is at the back of the queue.
func (q *PDeque) IsBack(e *Element) bool {
	i := mmheap.Max(&q.h)
	return q.h.elements[i] == e
}

// Update reorders the queue to reflect a change in e.Value that might cause e
// to occupy a different position within in the queue.
func (q *PDeque) Update(e *Element) {
	mmheap.Fix(&q.h, e.index)
}

// Remove removes e from the queue.
func (q *PDeque) Remove(e *Element) {
	mmheap.Remove(&q.h, e.index)
}

// Inverse returns an inverted "view" of q, such that q.Inverse().Pop() is
// equivalent to q.PopBack(), etc.
func (q *PDeque) Inverse() Queue {
	return inverse{q}
}

type inverse struct {
	*PDeque
}

func (q inverse) Peek() (e *Element, ok bool) {
	return q.PeekBack()
}

func (q inverse) Pop() (v interface{}, ok bool) {
	return q.PopBack()
}

func (q inverse) IsFront(e *Element) bool {
	return q.IsBack(e)
}
