package main

import "fmt"

func bad() {
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println(src, dst)
}

func correct() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println(src, dst)
}
