package kyu_test

import (
	. "github.com/dogmatiq/kyu/kyu"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("type PQueue", func() {
	var queue *PQueue[int]

	BeforeEach(func() {
		queue = NewPQueue[int]()
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
		It("places the smallest value at the front of the queue", func() {
			queue.Push(10)
			queue.Push(5)

			v, ok := queue.Peek()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(5))
		})

		It("places other values somewhere in the middle of the queue", func() {
			queue.Push(10)
			queue.Push(20)

			v, ok := queue.Peek()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(10))
		})
	})

	Describe("func Peek()", func() {
		It("returns the element with the smallest value", func() {
			queue.Push(2)
			queue.Push(-1)
			queue.Push(1)

			v, ok := queue.Peek()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(-1))
		})

		It("returns false if the queue is empty", func() {
			_, ok := queue.Peek()
			Expect(ok).To(BeFalse())
		})
	})

	Describe("func Pop()", func() {
		It("returns the smallest value", func() {
			queue.Push(2)
			queue.Push(-1)
			queue.Push(1)

			v, ok := queue.Pop()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(-1))
		})

		It("maintains the queue order", func() {
			queue.Push(2)
			queue.Push(-1)
			queue.Push(1)

			queue.Pop()

			v, ok := queue.Peek()
			Expect(ok).To(BeTrue())
			Expect(v).To(Equal(1))
		})

		It("returns false if the queue is empty", func() {
			_, ok := queue.Pop()
			Expect(ok).To(BeFalse())
		})
	})
})
