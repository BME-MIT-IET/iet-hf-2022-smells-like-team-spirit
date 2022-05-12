package main

import (
	"testing"
)

func TestGnomeSort(t *testing.T) {
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
		GnomeSort(tt.input)

		for i := 0; i < len(tt.input); i++ {
			if tt.input[i] != tt.output[i] {
				t.Errorf("%v != %v", tt.input, tt.output)
			}
		}
	}
}
