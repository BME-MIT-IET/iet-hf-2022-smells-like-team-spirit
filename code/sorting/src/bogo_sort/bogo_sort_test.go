//Part of Cosmos Project by OpenGenus Foundation
//Bogo Sort implementation on golang
//Written by Guilherme Lucas(guilhermeslucas)
package main

import (
	"fmt"
	"testing"
)

var inputs = [][]int{
	{1, 5, 8, 2, 6, 9},
	{1, 5, 8, 2, 6, 10, 1, 5, 8, 2, 6, 10},
	{1, 5, 8, 2, 6, 12, 1, 5, 8, 2, 6, 10, 1},
}

func BenchmarkBogoSort(b *testing.B) {
	for i := 0; i < len(inputs); i++ {
		b.Run(fmt.Sprintf("input_size_%d", len(inputs[i])), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				bogoSort(inputs[i])
			}
		})
	}
}
