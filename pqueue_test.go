package kyu_test

import (
	. "github.com/dogmatiq/kyu"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type PQueue", func() {
	var queue *PQueue

	BeforeEach(func() {
		queue = &PQueue{
			Less: func(a, b interface{}) bool {
				return a.(int) < b.(int)
			},
		}
	})

	Describe("func Len()", func() {
		It("returns zero when the queue is empty", func() {
			Expect(queue.Len()).To(Equal(0))
		})

		It("returns the number of elements in the queue", func() {
			queue.Push(1)
			queue.Push(2)

			Expect(queue.Len()).To(Equal(2))
		})
	})

	Describe("func Push()", func() {
		It("returns an element", func() {
			e := queue.Push(1)
			Expect(e).NotTo(BeNil())
			Expect(e.Value).To(Equal(1))
		})

		It("places the smallest value at the front of the queue", func() {
			queue.Push(10)

			e := queue.Push(5)
			Expect(queue.IsFront(e)).To(BeTrue())
		})

		It("places other values somewhere in the middle of the queue", func() {
			queue.Push(10)

			e := queue.Push(20)
			Expect(queue.IsFront(e)).To(BeFalse())
		})
	})

	Describe("func Peek()", func() {
		It("returns the element with the smallest value", func() {
			queue.Push(2)
			queue.Push(0)
			queue.Push(1)

			e, ok := queue.Peek()
			Expect(ok).To(BeTrue())
			Expect(e.Value).To(Equal(0))
		})

		It("returns false if the queue is empty", func() {
			_, ok := queue.Peek()
			Expect(ok).To(BeFalse())
		})
	})

	Describe("func Pop()", func() {
		It("returns the smallest value", func() {
			queue.Push(2)
			queue.Push(0)
			queue.Push(1)

			v, ok := queue.Pop()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(0))
		})

		It("maintains the queue order", func() {
			queue.Push(2)
			queue.Push(0)
			queue.Push(1)

			queue.Pop()

			e, ok := queue.Peek()
			Expect(ok).To(BeTrue())
			Expect(e.Value).To(Equal(1))
		})

		It("returns false if the queue is empty", func() {
			_, ok := queue.Pop()
			Expect(ok).To(BeFalse())
		})
	})

	Describe("func Update()", func() {
		It("repairs the queue order", func() {
			queue.Push(10)
			queue.Push(20)
			e := queue.Push(30)

			e.Value = 5
			queue.Update(e)
			Expect(queue.IsFront(e)).To(BeTrue())
		})
	})

	Describe("func Remove()", func() {
		It("removes the element from the queue", func() {
			queue.Push(10)
			e := queue.Push(20)
			queue.Push(30)

			queue.Remove(e)

			v, ok := queue.Pop()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(10))

			v, ok = queue.Pop()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(30))
		})
	})
})
