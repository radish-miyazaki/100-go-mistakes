package json

import (
	"encoding/json"
	"fmt"
)

type customer struct {
	ID         string
	Operations []float32
}

func main() {
	var s1 []float32
	customer1 := customer{
		ID:         "foo",
		Operations: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b))

	s2 := []float32{}
	customer2 := customer{
		ID:         "bar",
		Operations: s2,
	}
	b, _ = json.Marshal(customer2)
	fmt.Println(string(b))
}
