package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const NUMS_DICT = "123456789"

func main() {
	calibrationSum := 0

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
		fmt.Println(line)
		line = replaceNumbers(line)

		i, j := 0, len(line)-1
		n1, n2 := "", ""
		for i <= j && (n1 == "" || n2 == "") {
			if strings.Contains(NUMS_DICT, string(line[i])) {
				n1 = string(line[i])
			} else {
				i++
			}
			if strings.Contains(NUMS_DICT, string(line[j])) {
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
		fmt.Println(calibrationVal)
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

func replaceNumbers(line string) string {
	numsWords := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	i, j := 0, 0
	line += " "
	newLine := ""
	for i <= len(line)-1 && j <= len(line)-1 {
		for k, v := range numsWords {
			if strings.Contains(line[i:j+1], k) {
				newLine = newLine + v
				i = j
			}
			if strings.Contains(NUMS_DICT, string(line[i])) {
				newLine = newLine + string(line[i])
				i++
			}
			if strings.Contains(NUMS_DICT, string(line[j])) {
				newLine = newLine + string(line[j])
				i = j + 1
				j++
			}
		}
		j++
	}

	return newLine
}
