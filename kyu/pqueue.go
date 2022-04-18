package kyu

import "golang.org/x/exp/constraints"

// PQueue is a priority queue.
//
// It supports efficient inspection and removal of elements at the front of the
// queue. The "lowest" elements appear towards the front of the queue.
type PQueue[T any] struct {
	// Less is a function that determines the order of elements in the queue.
	//
	// If it returns true, a appears in the queue before b.
	Less func(a, b T) bool

	// heap is a min-heap containing the elements in the queue.
	heap []T
}

// NewPQueue returns a new priority queue containing elements of type T.
func NewPQueue[T constraints.Ordered]() *PQueue[T] {
	return &PQueue[T]{
		Less: func(a, b T) bool {
			return a < b
		},
	}
}

// NewReversePQueue returns a new priority queue containing elements of type T
// that places the "highest" elements at the front of the queue.
func NewReversePQueue[T constraints.Ordered]() *PQueue[T] {
	return &PQueue[T]{
		Less: func(a, b T) bool {
			return a > b
		},
	}
}

// Len returns the number of elements in the queue.
func (q *PQueue[T]) Len() int {
	return len(q.heap)
}

// Push adds a new element to the queue.
func (q *PQueue[T]) Push(v T) {
	// append the new element to the end of the slice
	n := len(q.heap)
	q.heap = append(q.heap, v)

	// move the new element upwards in the heap until it occupies an appropriate
	// node
	for n > 0 {
		parent := (n - 1) / 2

		if !q.less(n, parent) {
			// if n >= parent then the min-heap is intact
			break
		}

		// otherwise, n < parent and it belongs above it in the heap
		q.swap(parent, n)

		// continue at the next layer above
		n = parent
	}
}

// Pop removes the element at the head of the queue and returns it.
//
// If the queue is empty, ok is false and v is undefined.
func (q *PQueue[T]) Pop() (v T, ok bool) {
	n := len(q.heap) - 1

	if n == -1 {
		return v, false
	}

	// grab the element at the head of the queue, then move the last element to
	// the top.
	v = q.heap[0]
	q.heap[0] = q.heap[n]

	// clear the last element in the underlying array before shrinking the
	// slice. This allows garbage collection of the element when T is a pointer.
	var zero T
	q.heap[n] = zero
	q.heap = q.heap[:n]

	// move the element that's now at the top of the tree back down to an
	// appropriate position.
	parent := 0

	for {
		left := (2 * parent) + 1

		if left < 0 || left >= n {
			// left < 0 == integer overflow
			break
		}

		lesser := left
		if right := left + 1; right < n && q.less(right, left) {
			lesser = right
		}

		if !q.less(lesser, parent) {
			break
		}

		q.swap(parent, lesser)
		parent = lesser
	}

	return v, true
}

// Peek returns the element at the head of the queue without removing it from
// the queue.
//
// If the queue is empty, ok is false and v is undefined.
func (q *PQueue[T]) Peek() (v T, ok bool) {
	if len(q.heap) == 0 {
		return v, false
	}

	return q.heap[0], true
}

// less returns true if the element at i is less than the element at j.
func (q *PQueue[T]) less(i, j int) bool {
	return q.Less(
		q.heap[i],
		q.heap[j],
	)
}

// swap swaps the elements at the given indices.
func (q *PQueue[T]) swap(i, j int) {
	q.heap[i], q.heap[j] = q.heap[j], q.heap[i]
}
