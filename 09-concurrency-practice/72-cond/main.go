package main

import (
	"fmt"
	"sync"
	"time"
)

type Donation1 struct {
	mu      sync.RWMutex
	balance int
}

func listing1() {
	donation := &Donation1{}

	// リスナゴルーチン
	f := func(goal int) {
		donation.mu.RLock()
		for donation.balance < goal {
			donation.mu.RUnlock()
			donation.mu.RLock()
		}
		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.mu.RUnlock()
	}
	go f(10)
	go f(15)

	// 更新ゴルーチン
	go func() {
		for {
			time.Sleep(time.Second)
			donation.mu.Lock()
			donation.balance++
			donation.mu.Unlock()
		}
	}()
}

type Donation2 struct {
	ch      chan int
	balance int
}

func listing2() {
	donation := &Donation2{ch: make(chan int)}

	f := func(goal int) {
		for balance := range donation.ch {
			if balance >= goal {
				fmt.Printf("$%d goal reached\n", balance)
				return
			}
		}
	}

	go f(10)
	go f(15)

	for {
		time.Sleep(time.Second)
		donation.balance++
		donation.ch <- donation.balance
	}
}

type Donation3 struct {
	cond    *sync.Cond
	balance int
}

func listing3() {
	donation := &Donation3{cond: sync.NewCond(&sync.Mutex{})}

	f := func(goal int) {
		donation.cond.L.Lock()
		for donation.balance < goal {
			donation.cond.Wait()
		}

		fmt.Printf("$%d goal reached\n", donation.balance)
		donation.cond.L.Unlock()
	}
	go f(10)
	go f(15)

	for {
		time.Sleep(time.Second)
		donation.cond.L.Lock()
		donation.balance++
		donation.cond.Broadcast()
		donation.cond.L.Unlock()
	}
}

func main() {
	//listing1()
	//listing2()
	listing3()

	time.Sleep(20 * time.Second)
}
