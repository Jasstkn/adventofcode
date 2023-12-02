package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID    int
	Turns []Turn
}

type Turn struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	lines, err := readLinesFromFile("input.txt")
	if err != nil {
		slog.Error(err.Error())
		return
	}

	games := []Game{}
	for _, line := range lines {
		// skip empty lines
		if line == "" {
			continue
		}
		games = append(games, parseGame(line))
	}

	sum := 0
	blueprint := Turn{
		Red:   12,
		Green: 13,
		Blue:  14,
	}
	for g := range games {
		allTurnsMeetCondition := true
		for _, t := range games[g].Turns {
			if t.Red > blueprint.Red || t.Green > blueprint.Green || t.Blue > blueprint.Blue {
				allTurnsMeetCondition = false
				break
			}
		}

		if allTurnsMeetCondition {
			sum += games[g].ID
		}
	}
	fmt.Println(sum)
}

func parseGame(line string) Game {
	g := Game{}

	idStr := strings.Split(strings.Split(line, ":")[0], " ")[1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		slog.Error(err.Error())
		return g
	}
	g.ID = id

	g.Turns = []Turn{}
	turns := strings.Split(strings.Split(line, ":")[1], ";")
	for _, t := range turns {
		balls := strings.Split(t, ",")
		turn := Turn{}
		for _, ball := range balls {
			ball = strings.TrimSpace(ball)
			n, kind := strings.Fields(ball)[0], strings.Fields(ball)[1]
			switch kind {
			case "red":
				turn.Red, _ = strconv.Atoi(n)
			case "green":
				turn.Green, _ = strconv.Atoi(n)
			case "blue":
				turn.Blue, _ = strconv.Atoi(n)
			}
		}
		g.Turns = append(g.Turns, turn)
	}

	return g
}

func readLinesFromFile(filename string) ([]string, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(f), "\n")
	return lines, nil
}
