package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	if err := listing1(); err != nil {
		panic(err)
	}
}

func listing1() error {
	b := getMessage()
	var m map[string]any
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	fmt.Printf("%T\n", m["id"])
	return nil
}

func getMessage() []byte {
	return []byte(`{"id": 32, "name": "foo"}`)
}
