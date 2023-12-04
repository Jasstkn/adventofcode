package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	defer f.Close()

	sets := map[int]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")

		cardStr := strings.TrimSpace(strings.Replace(parts[0], "Card", "", -1))

		numersPart := strings.Split(parts[1], "|")
		winningCards, myCards := strings.Fields(numersPart[0]), strings.Fields(numersPart[1])

		wins := 0
		for _, c := range myCards {
			if isWinningCard(c, winningCards) {
				wins++
			}
		}
		card, _ := strconv.Atoi(cardStr)

		if _, ok := sets[card]; ok {
			for i := 0; i <= sets[card]; i++ {
				for i := card + 1; i <= card+wins; i++ {
					sets[i] += 1
				}
			}
			sets[card] += 1
		} else {
			for i := card; i <= card+wins; i++ {
				sets[i] += 1
			}
		}
	}

	points := 0
	for _, v := range sets {
		points += v
	}

	fmt.Println("total points: ", points)
}

func isWinningCard(card string, winningCards []string) bool {
	for _, winningCard := range winningCards {
		if winningCard == card {
			return true
		}
	}
	return false
}
