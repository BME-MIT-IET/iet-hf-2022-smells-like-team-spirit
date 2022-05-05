package main

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		input  [8]int
		output [8]int
	}{
		{[8]int{1, 2, 3, 4, 5, 6, 7, 8}, [8]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[8]int{8, 7, 6, 5, 4, 3, 2, 1}, [8]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[8]int{8, 7, 6, 5, 4, 4, 2, 1}, [8]int{1, 2, 4, 4, 5, 6, 7, 8}},
		{[8]int{6, 5, 4, 3, 2, 1, 0, -1}, [8]int{-1, 0, 1, 2, 3, 4, 5, 6}},
		{[8]int{}, [8]int{}},
	}

test:
	for _, tt := range tests {
		sorted := selectionSort(tt.input)

		for i := 0; i < len(sorted); i++ {
			if sorted[i] != tt.output[i] {
				t.Errorf("%v != %v", sorted, tt.output)
				continue test
			}
		}
	}
}
