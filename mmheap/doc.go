// Package mmheap provides a drop-in replacement for the container/heap package
// that provides min-max heap semantics.
//
// Any existing type that implements heap.Interface can be used with mmheap.
//
// min-max heaps provide efficient removal of both the minimum and maximum
// elements. The data structure is described in detail here:
//
// 	http://www.cs.otago.ac.nz/staffpriv/mike/Papers/MinMaxHeaps/MinMaxHeaps.pdf
package mmheap
