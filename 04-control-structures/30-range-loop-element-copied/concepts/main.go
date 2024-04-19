package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Printf("index=%d, value=%d\n", i, v)
	}

	for _, v := range s {
		fmt.Printf("value=%s\n", v)
	}
}
