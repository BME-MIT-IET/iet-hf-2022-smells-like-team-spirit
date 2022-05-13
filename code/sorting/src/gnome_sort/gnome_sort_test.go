package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time")

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

var inputs = [][]int{
	{1, 6, 2, 4, 9, 0, 5, 3, 7, 8},
	{19, 4, 14, 9, 10, 16, 3, 2, 15, 1, 13, 7, 20, 18, 17, 11, 5, 12, 6},
	{68, 20, 28, 100, 32, 50, 36, 3, 28, 16, 81, 45, 62, 82, 62, 73, 73, 52, 17, 11, 18, 3, 32, 63, 43, 73, 32, 24, 57, 82},
}

func generateInputs() {
	rand.Seed(time.Now().Unix())
	inputs = append(inputs, rand.Perm(1000))
	inputs = append(inputs, rand.Perm(10000))
	inputs = append(inputs, rand.Perm(100000))
	inputs = append(inputs, rand.Perm(250000))
	//inputs = append(inputs, rand.Perm(1000000)) // benchmark took almost 9 minutes for this one
	//inputs = append(inputs, rand.Perm(10000000)) // generating the numbers take some time, use this with caution // haven't used this, but I'm sure it would take too much time for the benchmark
	//inputs = append(inputs, rand.Perm(100000000)) // generating the numbers take some time, use this with caution
}

func BenchmarkGnomeSort(b *testing.B) {
	generateInputs()
	b.ResetTimer()
	for i := 0; i < len(inputs); i++ {
		b.Run(fmt.Sprintf("input_size=%d;", len(inputs[i])), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				GnomeSort(inputs[i])
			}
		})
	}
}
