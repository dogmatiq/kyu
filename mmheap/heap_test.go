package mmheap_test

import (
	"sort"

	. "github.com/dogmatiq/kyu/mmheap"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Context("min/max heap", func() {
	var (
		values  []int
		subject intHeap
	)

	BeforeEach(func() {
		// Prepare some values to use within the heap.
		//
		// We choose 10 values for the purpose of testing, as when arranged in a
		// heap this leaves some nodes with 2 children, one node with a single
		// child, and others with 0 children.
		//
		// The values are sorted so we can predict the order they *should*
		// appear when popping from the heap.
		values = []int{5, 10, 15, 20, 25, 30, 35, 40, 45, 50}

		subject = make(intHeap, len(values))
		copy(subject, values)

		Init(&subject)
	})

	Describe("func Push()", func() {
		It("places a smallest element at the front", func() {
			v := values[0] - 1 // make a value smaller than any other
			Push(&subject, v)
			Expect(subject[0]).To(Equal(v))
		})

		It("places the largest element according to Max()", func() {
			v := values[len(values)-1] + 1 // make a value larger than any other
			Push(&subject, v)
			i := Max(&subject)
			Expect(subject[i]).To(Equal(v))
		})

		It("maintains the heap structure", func() {
			subject = intHeap{}

			for _, v := range values {
				Push(&subject, v)
			}

			for _, x := range values {
				v := Pop(&subject)
				Expect(v).To(Equal(x))
			}
		})
	})

	Describe("func Pop()", func() {
		It("removes the elements from smallest to largest", func() {
			for _, x := range values {
				v := Pop(&subject)
				Expect(v).To(Equal(x))
			}
		})
	})

	Describe("func PopMax()", func() {
		It("removes the elements from largest to smallest", func() {
			for i := len(values) - 1; i >= 0; i-- {
				v := PopMax(&subject)
				Expect(v).To(Equal(values[i]))
			}
		})
	})

	Describe("func Remove()", func() {
		entries := []TableEntry{
			Entry("when the index is 0", 0),
			Entry("when index is on a min level", 1),
			Entry("when index is on a max level", 8),
		}

		DescribeTable(
			"returns the element at the given index",
			func(i int) {
				removed := subject[i]
				v := Remove(&subject, i)
				Expect(v).To(Equal(removed))
			},
			entries...,
		)

		DescribeTable(
			"removes element at the given index",
			func(i int) {
				removed := subject[i]
				Remove(&subject, i)

				for i, x := range values {
					if x == removed {
						values = append(values[:i], values[i+1:]...)
						break
					}
				}

				for _, x := range values {
					v := Pop(&subject)
					Expect(v).To(Equal(x))
				}
			},
			entries...,
		)
	})

	Describe("func Fix()", func() {
		DescribeTable(
			"maintains the heap structure",
			func(i, delta int) {
				changed := subject[i]
				subject[i] = changed + delta
				Fix(&subject, i)

				for i, x := range values {
					if x == changed {
						values[i] = changed + delta
					}
				}

				sort.Slice(
					values,
					func(i, j int) bool {
						return values[i] < values[j]
					},
				)

				for _, x := range values {
					v := Pop(&subject)
					Expect(v).To(Equal(x))
				}
			},
			Entry("when an element at index 0 gets larger", 0, +11),
			Entry("when an element on a min-level gets larger", 1, +11),
			Entry("when an element on a max-level gets larger", 8, +11),
			Entry("when an element on a min-level gets smaller", 1, -6),
			Entry("when an element on a max-level gets smaller", 8, -11),
		)
	})

	Describe("func Max()", func() {
		DescribeTable(
			"it returns the index of the maximum element",
			func(h intHeap, expect int) {
				i := Max(&h)
				Expect(i).To(Equal(expect))
			},
			Entry("when the heap is empty", intHeap{}, -1),
			Entry("when there is only a single element", intHeap{5}, 0),
			Entry("when there is only a left-hand child", intHeap{5, 10}, 1),
			Entry("when the left-hand child is greater", intHeap{5, 20, 10}, 1),
			Entry("when the right-hand child is greater", intHeap{5, 10, 20}, 2),
		)
	})
})

// intHeap is an implementation of heap.Interface for testing the min/max heap
// implementation.
type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	index := len(*h) - 1
	e := (*h)[index]
	*h = (*h)[:index]
	return e
}
