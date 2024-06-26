package main

import (
	"fmt"
)

func merge1(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for v := range ch1 {
			ch <- v
		}

		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func merge2(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}

		close(ch)
	}()

	return ch
}

func merge3(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case v, open := <-ch1:
				if !open {
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}

			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}
	}()

	return ch
}

func merge4(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				ch <- v
			}
		}

		close(ch)
	}()

	return ch
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	merged := merge4(ch1, ch2)

	go func() {
		for i := 0; i < 10; i++ {
			if i == 8 {
				close(ch2)
				continue
			}

			if i > 8 {
				ch1 <- i
			} else {
				ch2 <- i
			}

		}
		close(ch1)
	}()

	for v := range merged {
		fmt.Println(v)
	}
}
