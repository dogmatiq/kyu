package kyu

// pheap is an implementation of heap.Interface that stores *Element. It is the
// underlying storage for the priority queue implementations.
type pheap struct {
	less     func(a, b interface{}) bool
	elements []*Element
}

func (h *pheap) Len() int {
	return len(h.elements)
}

func (h *pheap) Less(i, j int) bool {
	return h.less(
		h.elements[i].Value,
		h.elements[j].Value,
	)
}

func (h *pheap) Swap(i, j int) {
	wasI := h.elements[i]
	wasJ := h.elements[j]

	h.elements[i] = wasJ
	wasJ.index = i

	h.elements[j] = wasI
	wasI.index = j
}

func (h *pheap) Push(x interface{}) {
	e := x.(*Element)
	e.index = len(h.elements)
	h.elements = append(h.elements, e)
}

func (h *pheap) Pop() interface{} {
	i := len(h.elements) - 1

	e := h.elements[i]
	e.index = -1

	h.elements[i] = nil
	h.elements = h.elements[:i]

	return e
}
