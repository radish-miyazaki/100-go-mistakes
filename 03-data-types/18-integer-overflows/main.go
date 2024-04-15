package main

import "math"

func Inc32(counter int32) int32 {
	if counter == math.MaxInt32 {
		panic("int32 overflow")
	}

	return counter + 1
}

func IncInt(counter int) int {
	if counter == math.MaxInt {
		panic("int overflow")
	}

	return counter + 1
}

func IncUint(counter uint) uint {
	if counter == math.MaxUint {
		panic("uint overflow")
	}

	return counter + 1
}

func AddInt(a, b int) int {
	if (b > 0 && a > math.MaxInt-b) || (b < 0 && a < math.MinInt-b) {
		panic("int overflow")
	}

	return a + b
}

func MultiplyInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	result := a * b
	if a == 1 || b == 1 {
		return result
	}

	if a == math.MinInt || b == math.MinInt {
		panic("integer overflow")
	}

	// 乗算が整数のオーバーフローを発生させているかチェック
	if result/b != a {
		panic("integer overflow")
	}

	return a * b
}
