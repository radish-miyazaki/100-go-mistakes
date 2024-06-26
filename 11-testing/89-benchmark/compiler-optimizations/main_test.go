package main

import "testing"

var global uint64

func BenchmarkPopcnt1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcnt(uint64(i))
	}
}

func BenchmarkPopcnt2(b *testing.B) {
	var v uint64
	for i := 0; i < b.N; i++ {
		v = popcnt(uint64(i))
	}
	global = v
}
