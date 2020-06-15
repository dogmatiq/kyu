package kyu

// Element is a container for a value on a queue.
//
// It couple's the element's value with queue-specific meta-data that allows
// manipulation of the element while it is on the queue.
type Element struct {
	// Value is the value associated with the element.
	Value interface{}

	index int
}

// Queue is an interface for a queue.
//
// The interface makes the distinction between a "value", which is the
// user-provided value to be placed on the queue, and an "element", which
// represents a value that has been placed on a queue.
type Queue interface {
	// Len returns the number of elements in the queue.
	Len() int

	// Push adds a new value to the queue.
	//
	// It returns the element that contains that value.
	Push(v interface{}) *Element

	// Peek returns the element at the front of the queue without removing it
	// from the queue.
	//
	// If the queue is empty, e is nil and ok is false.
	Peek() (e *Element, ok bool)

	// Pop removes the element at the front of the queue and returns its value.
	//
	// If the queue is empty, v is nil and ok is false.
	Pop() (v interface{}, ok bool)

	// Contains returns true if e is in the queue.
	Contains(e *Element) bool

	// IsFront returns true if e is at the front of the queue.
	IsFront(e *Element) bool

	// Update reorders the queue to reflect a change in e.Value that might cause
	// e to occupy a different position within in the queue.
	Update(e *Element)

	// Remove removes e from the queue.
	Remove(e *Element)
}
