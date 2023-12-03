package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Entry struct {
	slice    string
	startIdx int
	endIdx   int
}

func main() {
	lines, _ := readLinesFromFile("input.txt")

	entries := [][]Entry{}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		lineEntries := []Entry{}
		numStr := ""
		for j, c := range line {
			if strings.Contains("0123456789", string(c)) {
				numStr += string(c)
			} else {
				if numStr != "" {
					lineEntries = append(lineEntries, Entry{slice: numStr, startIdx: j - len(numStr), endIdx: j - 1})
				}
				numStr = ""
			}
			if j == len(line)-1 && numStr != "" {
				lineEntries = append(lineEntries, Entry{slice: numStr, startIdx: j - len(numStr) + 1, endIdx: j})
			}
		}
		entries = append(entries, lineEntries)
	}

	sum := 0
	for i, line := range entries {
		if len(line) == 0 {
			continue
		}
		for _, entry := range line {
			prevLine, currLine, nextLine := i-1, i, i+1

			leftIdx, rightIdx := entry.startIdx-1, entry.endIdx+1

			check := false
			fmt.Println(entry, leftIdx, rightIdx)
			if leftIdx < 0 {
				leftIdx = 0
			}
			if prevLine >= 0 && leftIdx >= 0 && rightIdx <= len(lines[prevLine]) {
				moditifiedLine := lines[prevLine] + "."
				if containsSpecialChar(moditifiedLine[leftIdx : rightIdx+1]) {
					check = true
				}
				fmt.Println("previous")
				fmt.Println(entry)
				fmt.Println(moditifiedLine[leftIdx : rightIdx+1])
			}
			if !check && leftIdx >= 0 && rightIdx <= len(lines[currLine]) {
				moditifiedLine := lines[currLine] + "."
				if containsSpecialChar(moditifiedLine[leftIdx : rightIdx+1]) {
					check = true
				}
				fmt.Println("current")
				fmt.Println(entry)
				fmt.Println(moditifiedLine[leftIdx : rightIdx+1])
			}
			if !check && leftIdx >= 0 && rightIdx <= len(lines[nextLine]) {
				moditifiedLine := lines[nextLine] + "."
				if containsSpecialChar(moditifiedLine[leftIdx : rightIdx+1]) {
					check = true
				}
				fmt.Println("next")
				fmt.Println(entry)
				fmt.Println(moditifiedLine[leftIdx : rightIdx+1])
			}
			if check {
				num, _ := strconv.Atoi(entry.slice)
				fmt.Println(num)
				sum += num
			}
		}
	}

	fmt.Println(sum)
}

func containsSpecialChar(s string) bool {
	for _, c := range s {
		if strings.Contains("0123456789.", string(c)) {
			continue
		}
		return true
	}
	return false
}

func readLinesFromFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(f), "\n")
	return lines, nil
}
