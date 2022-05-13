package main

import (
	"fmt"
	"testing"
)

func TestBogoSort(t *testing.T) {
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
		sorted := bogoSort(tt.input)

		for i := 0; i < len(sorted); i++ {
			if sorted[i] != tt.output[i] {
				t.Errorf("%v != %v", sorted, tt.output)
			}
		}
	}
}

var inputs = [][]int{
	{1, 5, 8, 2, 6, 9},
	{1, 5, 8, 2, 6, 10, 1, 5, 8, 2, 6, 10},
	{1, 5, 8, 2, 6, 12, 1, 5, 8, 2, 6, 10, 1},
}

func BenchmarkBogoSort(b *testing.B) {
	for i := 0; i < len(inputs); i++ {
		b.Run(fmt.Sprintf("input_size=%d;", len(inputs[i])), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				bogoSort(inputs[i])
			}
		})
	}
}