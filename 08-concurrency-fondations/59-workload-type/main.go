package main

import (
	"io"
	"sync"
	"sync/atomic"
)

func read1(r io.Reader) (int, error) {
	count := 0
	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, err
		}
		count += task(b)
	}

	return count, nil
}

func read2(r io.Reader) (int, error) {
	var count int64
	wg := sync.WaitGroup{}
	var n = 10

	ch := make(chan []byte, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for b := range ch {
				v := task(b)
				atomic.AddInt64(&count, int64(v))
			}
		}()
	}

	for {
		b := make([]byte, 1024)
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			return 0, nil
		}
		ch <- b
	}

	close(ch)
	wg.Wait()

	return int(count), nil
}

func task(b []byte) int {
	return len(b)
}
