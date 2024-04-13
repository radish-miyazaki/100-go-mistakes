package main

import (
	"fmt"
	"sort"
)

func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

type customConstraint interface {
	~int | ~string
}

func getKeyWithCustomConstraint[K customConstraint, V any](m map[K]V) []K {
	return nil
}

type Node[T any] struct {
	Val  T
	next *Node[T]
}

func (n *Node[T]) Add(next *Node[T]) {
	n.next = next
}

type SliceFn[T any] struct {
	S       []T
	Compare func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.S)
}

func (s SliceFn[T]) Less(i, j int) bool {
	return s.Compare(s.S[i], s.S[j])
}

func (s SliceFn[T]) Swap(i, j int) {
	s.S[i], s.S[j] = s.S[j], s.S[i]
}

func main() {
	s := SliceFn[int]{
		S: []int{3, 2, 1},
		Compare: func(i int, j int) bool {
			return i < j
		},
	}

	sort.Sort(s)
	fmt.Println(s.S)
}
