package main

import "errors"

func f() {
	// ...

	notify()
}

func f2() {
	// ...

	_ = notify()
}

func notify() error {
	return errors.New("failed to notify")
}
