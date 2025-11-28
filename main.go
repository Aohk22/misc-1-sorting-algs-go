package main

import (
	"fmt"
)

/*
1, 4, 2, 6 start list.
1, 4, 2, 6 index 0, no smaller element ahead.
1, 2, 4, 6 index 1, 2 is smaller, move to current index.
1, 2, 4, 6 index 2, no smaller element ahead.
*/
func SelectionSort[T string | int](slice []T, reverse bool) {
	for i := 0; i < len(slice)-1; i += 1 {
		for j := i + 1; j < len(slice); j += 1 {
			if slice[i] > slice[j] {
				t := slice[i]
				slice[i] = slice[j]
				slice[j] = t
			}
		}
	}
}

/*
1, 4, 2, 6 start list.
1, 4, 2, 6 index 1: 1 automatically in list.
1, 4, 2, 6 index 2: add 4 to list, already sorted.
1, 2, 4, 6 index 3: add 2 to list, bump up 1 time (to before 4).
1, 2, 4, 6 index 4: add 6 to list, already sorted.
*/
func InsertionSort[T string | int](slice []T, reverse bool) {
	for i := 0; i < len(slice); i += 1 {
		for i > 0 && slice[i] < slice[i-1] {
			t := slice[i]
			slice[i] = slice[i-1]
			slice[i-1] = t
			i -= 1
		}
	}
}

/*
1, 4, 2, 6, 7 split first, then merge using recursion.
*/
func MergeSort[T string | int](slice []T, reverse bool) ([]T) {
	if len(slice) < 2 {
		return slice
	}
	// split
	sortedSlice1 := MergeSort(slice[0:len(slice)/2], false)
	sortedSlice2 := MergeSort(slice[len(slice)/2:], false)
	// merge the slices
	i := 0
	j := 0
	sorted := make([]T, 0, len(sortedSlice1) + len(sortedSlice2))
	for i < len(sortedSlice1) && j < len(sortedSlice2) {
		if sortedSlice1[i] < sortedSlice2[j] {
			sorted = append(sorted, sortedSlice1[i])
			i += 1
		} else {
			sorted = append(sorted, sortedSlice2[j])
			j += 1
		}
	}
	// append leftovers
	sorted = append(sorted, sortedSlice1[i:]...)
	sorted = append(sorted, sortedSlice2[j:]...)
	return sorted
}

func QuickSort[T string | int](slice []T, reverse bool) ([]T) {
	if len(slice) < 2 {
		return slice
	}
	indexForLarger := 0
	indexPivot := len(slice) - 1
	indexForSmaller := indexPivot - 1
	pivotValue := slice[indexPivot]

	if indexForLarger == indexForSmaller {
		if slice[indexForLarger] > slice[indexPivot] {
			t := slice[indexForLarger]
			slice[indexForLarger] = slice[indexPivot]
			slice[indexPivot] = t
		}
		return slice
	}
	for indexForLarger < indexForSmaller {
		if slice[indexForLarger] > slice[indexPivot] {
			for indexForLarger < indexForSmaller {
				if slice[indexForSmaller] < slice[indexPivot] {
					t := slice[indexForLarger]
					slice[indexForLarger] = slice[indexForSmaller]
					slice[indexForSmaller] = t
					indexForSmaller = indexPivot - 1
					break
				}
				indexForSmaller -= 1
			}
		}
		indexForLarger += 1
	}
	// array is now shorted around pivot, now sort it recursively with another pivot.
	sortedLeft := QuickSort(slice[:indexForSmaller], false)
	sortedRight := QuickSort(slice[indexForSmaller:indexPivot], false)
	result := make([]T, 0, len(slice))
	result = append(result, sortedLeft...)
	result = append(result, pivotValue)
	result = append(result, sortedRight...)
	return result
}

func main() {
	var a []int = []int{1, 4, 2, 6, 3, 9, 20, 7}
	fmt.Printf("Before: %v\n", a)
	a = MergeSort(a, false)
	fmt.Printf("After : %v\n", a)
}
