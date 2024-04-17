package main

func handleOperations(id string) {
	operations := getOprations(id)
	if operations != nil {
		handle(operations)
	}
}

func handleOperations2(id string) {
	operations := getOprations(id)
	if len(operations) != 0 {
		handle(operations)
	}
}

func getOprations(id string) []float32 {
	operations := make([]float32, 0)

	if id == "" {
		return operations
	}

	return operations
}

func handle(operations []float32) {}
