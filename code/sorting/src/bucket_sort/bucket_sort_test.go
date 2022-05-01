/* Part of Cosmos by OpenGenus Foundation */
package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

var inputs []struct {
	Array      []float64
	BucketSize int
}

func generateInputs() {
	rand.Seed(time.Now().Unix())
	inputs = append(inputs, struct {
		Array      []float64
		BucketSize int
	}{
		Array: randFloats(float64(0), float64(rand.Intn(100)), rand.Intn(100)), BucketSize: Max(1, rand.Intn(10)),
	})
	inputs = append(inputs, struct {
		Array      []float64
		BucketSize int
	}{
		Array: randFloats(float64(0), float64(rand.Intn(1000)), rand.Intn(1000)), BucketSize: Max(1, rand.Intn(30)),
	})
	inputs = append(inputs, struct {
		Array      []float64
		BucketSize int
	}{
		Array: randFloats(float64(0), float64(rand.Intn(10000)), rand.Intn(10000)), BucketSize: Max(1, rand.Intn(50)),
	})
	inputs = append(inputs, struct {
		Array      []float64
		BucketSize int
	}{
		Array: randFloats(float64(0), float64(rand.Intn(100000)), rand.Intn(100000)), BucketSize: Max(1, rand.Intn(50)),
	})
}

func randFloats(min, max float64, n int) []float64 {
	res := make([]float64, n)
	for i := range res {
		res[i] = min + rand.Float64()*(max-min)
	}
	return res
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func BenchmarkBucketSort(b *testing.B) {
	generateInputs()
	b.ResetTimer()
	for i := 0; i < len(inputs); i++ {
		b.Run(fmt.Sprintf("input_size=%d;bucket_size=%d;", len(inputs[i].Array), inputs[i].BucketSize), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				bucketSort(inputs[i].Array, inputs[i].BucketSize)
			}
		})
	}
}
