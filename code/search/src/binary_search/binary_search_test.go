package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input  []int
		search int
		output int
	}{
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 6, -1},
	}

	for _, tt := range tests {
		index := binarySearch(tt.input, tt.search)

		if index != tt.output {
			t.Errorf("%v != %v", index, tt.output)
		}
	}
}
