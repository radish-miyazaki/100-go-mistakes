package main

import (
	"bufio"
	"io"
	"os"
)

func countEmptyLinesInFile(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// file の Close を遅延させる処理

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// ...
	}

	return 0, nil
}

func countEmptyLines(reader io.Reader) (int, error) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		// ...
	}

	return 0, nil
}
