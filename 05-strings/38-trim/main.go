package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.TrimRight("123oxo", "xo"))
	fmt.Println(strings.TrimRight("oxo123", "xo"))

	fmt.Println(strings.TrimSuffix("123oxo", "xo"))
	fmt.Println(strings.TrimPrefix("oxo123", "xo"))
}
