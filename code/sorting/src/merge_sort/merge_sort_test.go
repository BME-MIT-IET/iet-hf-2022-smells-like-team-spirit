package main

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 4, 2, 1}, []int{1, 2, 4, 4, 5}},
		{[]int{3, 2, 1, 0, -1}, []int{-1, 0, 1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		sorted := sort(tt.input)

		for i := 0; i < len(sorted); i++ {
			if sorted[i] != tt.output[i] {
				t.Errorf("%v != %v", sorted, tt.output)
			}
		}
	}
}
