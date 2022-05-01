/* Part of Cosmos by OpenGenus Foundation */
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var inputs = [][]int{
	{1, 6, 2, 4, 9, 0, 5, 3, 7, 8},
	{19, 4, 14, 9, 10, 16, 3, 2, 15, 1, 13, 7, 20, 18, 17, 11, 5, 12, 6},
	{68, 20, 28, 100, 32, 50, 36, 3, 28, 16, 81, 45, 62, 82, 62, 73, 73, 52, 17, 11, 18, 3, 32, 63, 43, 73, 32, 24, 57, 82},
}

func generateInputs() {
	rand.Seed(time.Now().Unix())
	inputs = append(inputs, rand.Perm(1000))
	inputs = append(inputs, rand.Perm(10000))
	//inputs = append(inputs, rand.Perm(100000)) // takes almost half a minute on my PC
	//inputs = append(inputs, rand.Perm(250000)) // takes more than two minutes on my PC
	//inputs = append(inputs, rand.Perm(1000000)) // benchmark timed out, don't use this!
	//inputs = append(inputs, rand.Perm(10000000)) // generating the numbers take some time, use this with caution // haven't used this, but I'm sure it would take too much time for the benchmark
	//inputs = append(inputs, rand.Perm(100000000)) // generating the numbers take some time, use this with caution
}

func BenchmarkCycleSort(b *testing.B) {
	generateInputs()
	b.ResetTimer()
	for i := 0; i < len(inputs); i++ {
		b.Run(fmt.Sprintf("input_size=%d;", len(inputs[i])), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				CycleSort(&inputs[i])
			}
		})
	}
}
