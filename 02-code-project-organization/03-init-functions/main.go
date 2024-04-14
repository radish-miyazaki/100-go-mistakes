package main

import (
	"fmt"

	"github.com/radish-miyazaki/100-go-mistakes/02-code-project-organization/03-init-functions/redis"
)

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}

func main() {
	err := redis.Store("foo", "bar")
	_ = err
}
