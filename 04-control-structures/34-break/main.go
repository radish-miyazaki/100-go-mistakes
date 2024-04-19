package main

import "fmt"

func main() {
	listing1()
	listing2()
}

func listing1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break
		}
	}
}

func listing2() {
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", i)

		switch i {
		default:
		case 2:
			break loop
		}

	}
}
