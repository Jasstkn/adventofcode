package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input_example.txt")
	defer f.Close()

	var points float64
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		cardN := strings.TrimSpace(parts[0])

		numersPart := strings.Split(parts[1], "|")
		winningCards, myCards := strings.Fields(numersPart[0]), strings.Fields(numersPart[1])
		fmt.Printf("Card number: %s, winning cards: %s, myCards: %s \n", cardN, winningCards, myCards)

		cnt := 0
		for _, card := range myCards {
			if isWinningCard(card, winningCards) {
				cnt++
			}
		}

		if cnt > 0 {
			fmt.Println("card: ", cardN, ",adding points: ", math.Pow(2, float64(cnt-1)))
			points += math.Pow(2, float64(cnt-1))
		}
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
