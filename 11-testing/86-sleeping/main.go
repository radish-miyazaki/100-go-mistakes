package main

type Handler struct {
	n         int
	publisher publisher
}

type publisher interface {
	Publish([]Foo)
}

type Foo struct{}

func (h Handler) getBestFoo(someInputs int) Foo {
	foos := getFoos(someInputs)
	best := foos[0]

	go func() {
		if len(foos) > h.n {
			foos = foos[:h.n]
		}
		h.publisher.Publish(foos)
	}()

	return best
}

func getFoos(_ int) []Foo {
	return make([]Foo, 100)
}
