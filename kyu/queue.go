package kyu

// Element is an element on a queue.
//
// It can be used to efficiently remove or reorder the element within the queue.
type Element[T any] struct {
	value T
	queue Queue[T]
	index int
}

type Queue[T any] interface {
	Len() int
	Push(T) Element[T]
	Pop() (T, bool)
	Peek() (Element[T], bool)

	IsFront(Element[T])
	Contains(Element[T])
	Remove(Element[T])
	Update(Element[T], T)
}
