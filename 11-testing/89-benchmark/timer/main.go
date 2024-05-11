package main

import "testing"

func BenchmarkFoo1(b *testing.B) {
	expensiveSetup()
	for i := 0; i < b.N; i++ {
		functionUnderTest()
	}
}

func BenchmarkFoo2(b *testing.B) {
	expensiveSetup()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		functionUnderTest()
	}
}

func functionUnderTest() {}

func expensiveSetup() {}
