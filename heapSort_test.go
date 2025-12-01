package main

import (
	"testing"
	"reflect"
)

type test[T int | string] struct {
	name string
	input []T
	want []T
}

func TestHeapSort(t *testing.T) {
	tests := []test[int] {
		{"basic", []int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			got := HeapSort(tt.input)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
