package main

import (
	"fmt"
	"sync"
)

func listing1() {
	s := []int{1, 2, 3}

	wg := sync.WaitGroup{}
	for _, i := range s {
		wg.Add(1)
		go func() {
			// INFO:
			//	~ v1.21 の場合、i に入る値はゴルーチンの実行時に決定されるため、予期せぬ結果が出力される
			// 	しかし、v1.22 以降では、イテレートごとに新しい変数が作成されるため、期待通りの結果が出力される
			defer wg.Done()
			fmt.Print(i)
		}()
	}

	wg.Wait()
}

// 以下、~ v1.21 での回避策

func listing2() {
	s := []int{1, 2, 3}

	wg := sync.WaitGroup{}
	for _, i := range s {
		val := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Print(val)
		}()
	}

	wg.Wait()
}

func listing3() {
	s := []int{1, 2, 3}

	wg := sync.WaitGroup{}
	for _, i := range s {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			fmt.Print(val)
		}(i)
	}

	wg.Wait()
}
