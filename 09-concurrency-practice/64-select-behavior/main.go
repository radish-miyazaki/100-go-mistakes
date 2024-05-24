package main

import (
	"fmt"
	"time"
)

func main() {
	messageCh := make(chan int)
	disconnectCh := make(chan struct{})

	go listing1(messageCh, disconnectCh)

	for i := 0; i < 3; i++ {
		go producer(i, messageCh)
	}
	disconnectCh <- struct{}{}

	time.Sleep(1 * time.Second)

	//for i := 0; i < 10; i++ {
	//	messageCh <- i
	//}
	//disconnectCh <- struct{}{}
	//time.Sleep(10 * time.Millisecond)
}

func listing1(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			fmt.Println("disconnection, return")
			return
		}
	}
}

func listing2(messageCh <-chan int, disconnectCh chan struct{}) {
	for {
		select {
		case v := <-messageCh:
			fmt.Println(v)
		case <-disconnectCh:
			for {
				select {
				case v := <-messageCh:
					fmt.Println(v)
				default:
					fmt.Println("disconnection, return")
					return
				}
			}
		}
	}
}

func producer(id int, messageCh chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Producer %d: %d\n", id, i)
		messageCh <- i
		time.Sleep(100 * time.Millisecond)
	}
}
