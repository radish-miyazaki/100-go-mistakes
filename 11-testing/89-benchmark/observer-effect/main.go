package main

func calculateSum1024(s [][1024]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}
	return sum
}

func calculateSum1025(s [][1025]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}
	return sum
}
