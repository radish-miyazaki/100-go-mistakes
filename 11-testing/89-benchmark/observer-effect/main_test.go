package main

import "testing"

const rows = 1000

var res int64

func BenchmarkCalculateSum1024_1(b *testing.B) {
	var sum int64

	s := createMatrix1024(rows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum1024(s)
	}
	res = sum
}

func BenchmarkCalculateSum1025_1(b *testing.B) {
	var sum int64

	s := createMatrix1025(rows)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sum = calculateSum1025(s)
	}
	res = sum
}

func BenchmarkCalculateSum1024_2(b *testing.B) {
	var sum int64

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := createMatrix1024(rows)
		b.StartTimer()
		sum = calculateSum1024(s)
	}
	res = sum
}

func BenchmarkCalculateSum1025_2(b *testing.B) {
	var sum int64

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := createMatrix1025(rows)
		b.StartTimer()
		sum = calculateSum1025(s)
	}
	res = sum
}

func createMatrix1024(r int) [][1024]int64 {
	return make([][1024]int64, r)
}

func createMatrix1025(r int) [][1025]int64 {
	return make([][1025]int64, r)
}
