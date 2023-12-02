package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readLinesFromFile("input.txt")
	if err != nil {
		return
	}

	sum := 0
	for _, line := range lines {
		// skip empty lines
		if line == "" {
			continue
		}
		sum += calculateWrappingPaper(line)
	}
	fmt.Println(sum)
}

func calculateWrappingPaper(input string) int {
	dimensions := strings.Split(strings.TrimSpace(input), "x")
	lStr, wStr, hStr := dimensions[0], dimensions[1], dimensions[2]
	l, _ := strconv.Atoi(lStr)
	w, _ := strconv.Atoi(wStr)
	h, _ := strconv.Atoi(hStr)

	return l*w*h + findPerimeter(find2Min(l, w, h))
}

func findPerimeter(a, b int) int {
	return 2*a + 2*b
}

func find2Min(a, b, c int) (int, int) {
	if a <= b && a <= c {
		return a, findMin(b, c)
	} else if b <= a && b <= c {
		return b, findMin(a, c)
	}
	return c, findMin(a, b)
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func readLinesFromFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(f), "\n")
	return lines, nil
}
