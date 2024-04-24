package main

import "sync"

func sequentialMergeSort(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2
	sequentialMergeSort(s[:middle])
	sequentialMergeSort(s[middle:])
	merge(s, middle)
}

func parallelMergeSortV1(s []int) {
	if len(s) <= 1 {
		return
	}

	middle := len(s) / 2

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		parallelMergeSortV1(s[:middle])
	}()

	go func() {
		defer wg.Done()
		parallelMergeSortV1(s[:middle])
	}()

	wg.Wait()
	merge(s, middle)
}

const MAX = 2048

func parallelMergeSortV2(s []int) {
	if len(s) <= 1 {
		return
	}

	if len(s) <= MAX {
		sequentialMergeSort(s)
	} else {
		middle := len(s) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			parallelMergeSortV2(s[:middle])
		}()

		go func() {
			defer wg.Done()
			parallelMergeSortV2(s[middle:])
		}()

		wg.Wait()
		merge(s, middle)
	}
}

func merge(s []int, middle int) {
	helper := make([]int, len(s))
	copy(helper, s)

	helperLeft := 0
	helperRight := middle
	current := 0
	high := len(s) - 1

	for helperLeft <= middle-1 && helperRight <= high {
		if helper[helperLeft] <= helper[helperRight] {
			s[current] = helper[helperLeft]
			helperLeft++
		} else {
			s[current] = helper[helperRight]
			helperRight++
		}
		current++
	}

	// 残った左側の要素をコピー
	for helperLeft <= middle-1 {
		s[current] = helper[helperLeft]
		current++
		helperLeft++
	}
}
