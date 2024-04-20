package main

import "fmt"

func main() {
	s := "hello"
	fmt.Println(len(s))

	s = "漢"
	fmt.Println(len(s))

	s := string([]byte{0xE6, 0xBC, 0xA2})
	fmt.Printf("%s\n", s)
}
