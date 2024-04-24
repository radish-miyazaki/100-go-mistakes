package main

import (
	"math/rand"
	"testing"
	"time"
)

var global []int

func Benchmark_sequentialMergeSort(b *testing.B) {
	var local []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandomElements()
		b.StartTimer()

		sequentialMergeSort(input)
		local = input
	}
	global = local
}

func Benchmark_parallelMergeSortV1(b *testing.B) {
	var local []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandomElements()
		b.StartTimer()

		parallelMergeSortV1(input)
		local = input
	}
	global = local
}

func Benchmark_parallelMergeSortV2(b *testing.B) {
	var local []int
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		input := getRandomElements()
		b.StartTimer()

		parallelMergeSortV2(input)
		local = input
	}
	global = local
}

func getRandomElements() []int {
	n := 10_000
	res := make([]int, n)
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	for i := 0; i < n; i++ {
		res[i] = rnd.Int()
	}

	return res
}
