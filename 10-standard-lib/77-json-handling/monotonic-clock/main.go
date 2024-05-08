package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func main() {
	err := listing1()
	if err != nil {
		fmt.Println(err)
	}

	err = listing2()
	if err != nil {
		fmt.Println(err)
	}
}

type Event struct {
	Time time.Time
}

func listing1() error {
	t := time.Now()
	event1 := Event{
		Time: t,
	}
	b, err := json.Marshal(event1)
	if err != nil {
		return err
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		return err
	}

	fmt.Println(event1 == event2)
	return nil
}

func listing2() error {
	t := time.Now()
	event1 := Event{
		Time: t.Truncate(0),
	}
	b, err := json.Marshal(event1)
	if err != nil {
		return err
	}

	var event2 Event
	err = json.Unmarshal(b, &event2)
	if err != nil {
		return err
	}

	t = time.Now().UTC()
	fmt.Println(t)

	fmt.Println(event1 == event2)
	return nil
}
