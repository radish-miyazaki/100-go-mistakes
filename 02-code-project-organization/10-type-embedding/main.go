package main

import (
	"io"
	"sync"
)

type InMem struct {
	mu sync.Mutex
	m  map[string]int
}

func New() *InMem {
	return &InMem{m: make(map[string]int)}
}

func (i *InMem) Get(key string) (int, bool) {
	i.mu.Lock()
	v, contains := i.m[key]
	i.mu.Unlock()

	return v, contains
}

type Logger struct {
	io.WriteCloser
}

func main() {
	l := Logger{}
	_, _ = l.Write([]byte("byte"))
	_ = l.Close()
}
