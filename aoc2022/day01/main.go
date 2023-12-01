package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Calories []int
}

func main() {
	fileContent, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	fileLen := countRune(fileContent, '\n')

	elves := make([]Elf, fileLen/2)
	calories := make([]int, fileLen/2)
	idx := 0
	lines := strings.Split(string(fileContent), "\n")
	for _, line := range lines {
		if line == "" {
			elves[idx].Calories = calories
			idx++
			calories = []int{}
			continue
		}
		caloriesInt, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		calories = append(calories, caloriesInt)
	}
	fmt.Println(findFattestNElves(elves, 3))
}

func countRune(s []byte, r rune) int {
	count := 0
	for _, c := range s {
		if rune(c) == r {
			count++
		}
	}
	return count
}

func findFattestNElves(elves []Elf, n int) int {
	var total, calories int
	elvesCalories := make([]int, len(elves))
	for _, elf := range elves {
		for _, c := range elf.Calories {
			calories += c
		}
		elvesCalories = append(elvesCalories, calories)
		calories = 0
	}
	sort.Ints(elvesCalories)
	for _, elf := range elvesCalories[len(elvesCalories)-n:] {
		total += elf
	}
	return total
}
