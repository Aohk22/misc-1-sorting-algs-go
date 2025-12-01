package main

/*
Description: 
Components:
build-max-heap: creates max heap from unsorted array.
heapify: build-max-heap but assume part of array is already sorted.

Process: build-max-heap -> heap-sort -> heapify -> heap-sort -> heapify etc.

Reference for pseudocode: https://youtu.be/2DmK_H7IdTo?si=FOo9MdHrOt34RD4U
*/

func BuildMaxHeap[T int | string](slice []T) {
	n := len(slice)
	startIdx := (n-1) / 2

	// this sorts the slice, we need to interate only once cause Heapify recurse will resolve
	// any inconsistency after the swap.
	for i := int(startIdx); i >= 0; i-- {
		Heapify(slice, i)
	}
}

func Heapify[T int | string](slice []T, i int) {
	var biggest int
	left := 2*i
	right := 2*i + 1
	n := len(slice)

	if left < n && slice[left] > slice[i] {
		biggest = left
	} else {
		biggest = i
	}

	if right < n && slice[right] > slice[biggest] {
		biggest = right
	}

	if (biggest != i) {
		t := slice[i]
		slice[i] = slice[biggest]
		slice[biggest] = t
		Heapify(slice, biggest)
	}
}

func HeapSort[T int | string](slice []T) ([]T) {
	BuildMaxHeap(slice)

	n := len(slice) - 1

	for i := n; i >= 0; i -= 1 {
		t := slice[0]
		slice[0] = slice[i]
		slice[i] = t

		Heapify(slice[:i], 0)
		// pass in with reduced length so Heapify doesn't evaluate the last element.
		// length shrinks, capacity stay the same.
	}
	return slice
}
