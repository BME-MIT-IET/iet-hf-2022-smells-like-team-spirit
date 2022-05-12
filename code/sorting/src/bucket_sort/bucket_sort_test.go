package main

import (
	"math"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		input  []float64
		output []float64
	}{
		{[]float64{1.0, 2.0, 3.0, 4.0, 5.0}, []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
		{[]float64{5.0, 4.0, 3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0, 4.0, 5.0}},
		{[]float64{5.0, 4.0, 4.0, 2.0, 1.0}, []float64{1.0, 2.0, 4.0, 4.0, 5.0}},
		{[]float64{3.0, 2.0, 1.0, 0.0, -1.0}, []float64{-1.0, 0.0, 1.0, 2.0, 3.0}},
		{[]float64{}, []float64{}},
		{[]float64{math.NaN(), 4.0, 3.0, 2.0, 1.0}, []float64{1.0, 2.0, 3.0, 4.0, math.NaN()}},
	}

	for _, tt := range tests {
		sorted := bucketSort(tt.input, 3)

		for i := 0; i < len(tt.input); i++ {
			if sorted[i] != tt.output[i] {
				t.Errorf("%v != %v", sorted, tt.output)
			}
		}
	}
}
