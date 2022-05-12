package main

import (
	"testing"
)

func TestTreeSort(t *testing.T) {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 4, 4, 2, 1}, []int{1, 2, 4, 4, 5}},
		{[]int{3, 2, 1, 0, -1}, []int{-1, 0, 1, 2, 3}},
		{[]int{}, nil},
	}

	for _, tt := range tests {
		tree := binaryTreeSort(tt.input)

		i := 0

		if tree == nil {
			if tt.output != nil {
				t.Errorf("Output is nil when it should be %v", tt.output)
			}

			continue
		}

		tree.forEach(func(v int) {
			if v != tt.output[i] {
				t.Errorf("%v is not equal to the sorted array", tt.output)
			}
			i++
		})
	}
}
