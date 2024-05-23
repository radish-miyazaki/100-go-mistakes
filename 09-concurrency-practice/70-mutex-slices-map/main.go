package main

import (
	"strconv"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func (c *Cache) AddBalance(id string, balance float64) {
	c.mu.Lock()
	c.balances[id] = balance
	c.mu.Unlock()
}

func (c *Cache) AverageBalance1() float64 {
	c.mu.RLock()
	balances := c.balances
	c.mu.RUnlock()

	sum := 0.
	for _, balance := range balances {
		sum += balance
	}

	return sum / float64(len(balances))
}

func (c *Cache) AverageBalance2() float64 {
	c.mu.RLock()
	defer c.mu.RUnlock()

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}

	return sum / float64(len(c.balances))
}

func (c *Cache) AverageBalance3() float64 {
	c.mu.RLock()
	m := make(map[string]float64, len(c.balances))
	for k, v := range c.balances {
		m[k] = v
	}
	defer c.mu.RUnlock()

	if len(c.balances) == 0 {
		return 0
	}

	sum := 0.
	for _, balance := range c.balances {
		sum += balance
	}

	return sum / float64(len(c.balances))
}

func main() {
	c := Cache{balances: map[string]float64{}}
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			c.AddBalance(strconv.Itoa(i), float64(i))
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			//c.AverageBalance1()
			//c.AverageBalance2()
			c.AverageBalance3()
			time.Sleep(time.Millisecond)
		}
	}()

	wg.Wait()
}
