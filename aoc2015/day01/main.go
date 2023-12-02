package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	f, err := readLinesFromFile("input.txt")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	floor := 0
	for _, ch := range f {
		if ch == '(' {
			floor++
		} else if ch == ')' {
			floor--
		}
		
	}
	fmt.Println(floor)
}

func readLinesFromFile(filename string) (string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(f), nil
}
