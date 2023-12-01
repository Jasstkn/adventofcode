package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	calibrationSum := 0
	numsDict := "0123456789"

	lines, err := readLinesFromFile("input.txt")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	for _, line := range lines {
		// skip empty lines
		if line == "" {
			continue
		}

		i, j := 0, len(line)-1
		n1, n2 := "", ""
		for i <= j && (n1 == "" || n2 == "") {
			if strings.Contains(numsDict, string(line[i])) {
				n1 = string(line[i])
			} else {
				i++
			}
			if strings.Contains(numsDict, string(line[j])) {
				n2 = string(line[j])
			} else {
				j--
			}
		}

		n1, n2 = checkEmptyNumbers(n1, n2)
		calibrationVal, err := strconv.Atoi(n1 + n2)
		if err != nil {
			slog.Error(line)
		}

		calibrationSum += calibrationVal
	}

	fmt.Println(calibrationSum)
}

func readLinesFromFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(f), "\n")
	return lines, nil
}

func checkEmptyNumbers(n1, n2 string) (string, string) {
	if n1 == "" {
		return n2, n2
	}
	if n2 == "" {
		return n1, n1
	}
	return n1, n2
}
