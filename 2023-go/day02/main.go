package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Number  int
	Reveals []Subset
}

type Subset struct {
	Red   int
	Green int
	Blue  int
}

// Full problem text at: https://adventofcode.com/2023/day/2

func main() {
	file, err := os.Open("day02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	games := []Game{}

	for scanner.Scan() {
		line := scanner.Text()
		game := ParseLine(line)
		games = append(games, game)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(games, 12, 13, 14))
	fmt.Println("Second Part: ", SecondPart(games))
}

func FirstPart(games []Game, totalRed int, totalGreen int, totalBlue int) int {
	gameIdSum := 0

	for _, game := range games {
		isGameValid := true
		for _, reveal := range game.Reveals {
			if reveal.Red > totalRed || reveal.Green > totalGreen || reveal.Blue > totalBlue {
				isGameValid = false
				break
			}
		}

		if isGameValid {
			gameIdSum += game.Number
		}
	}

	return gameIdSum
}

func SecondPart(games []Game) int {
	powerSum := 0

	for _, game := range games {
		powerSum += gamePower(game)
	}

	return powerSum
}

func minimumCubes(game Game) (minRed int, minGreen int, minBlue int) {
	maxRed := 0
	maxGreen := 0
	maxBlue := 0

	for _, reveal := range game.Reveals {
		if reveal.Red > maxRed {
			maxRed = reveal.Red
		}

		if reveal.Green > maxGreen {
			maxGreen = reveal.Green
		}

		if reveal.Blue > maxBlue {
			maxBlue = reveal.Blue
		}
	}

	return maxRed, maxGreen, maxBlue
}

func gamePower(game Game) int {
	minRed, minGreen, minBlue := minimumCubes(game)
	return minRed * minGreen * minBlue
}

func ParseLine(line string) Game {
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 3 {
		panic("Unexpected line format")
	}

	gameNumber, _ := strconv.Atoi(matches[1])
	contents := matches[2]

	revealsRaw := strings.Split(contents, "; ")
	var reveals []Subset

	for _, reveal := range revealsRaw {
		reveals = append(reveals, ParseSubset(reveal))
	}

	return Game{
		Number:  gameNumber,
		Reveals: reveals,
	}
}

func ParseSubset(input string) Subset {
	var subset Subset
	parts := strings.Split(input, ", ")

	re := regexp.MustCompile(`(\d+)\s*(\w+)`)

	for _, part := range parts {
		matches := re.FindStringSubmatch(part)
		if len(matches) < 3 {
			panic(fmt.Sprintf("invalid part: %s", part))
		}

		number, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(fmt.Sprintf("invalid number in part: %s", part))
		}

		switch matches[2] {
		case "red":
			subset.Red = number
		case "green":
			subset.Green = number
		case "blue":
			subset.Blue = number
		default:
			panic(fmt.Sprintf("unknown color: %s", matches[2]))
		}
	}

	return subset
}

func PrintGame(game Game) {
	fmt.Printf("Game Number: %d\n", game.Number)
	for _, reveal := range game.Reveals {
		fmt.Printf("Reveal - Red: %d, Green: %d, Blue: %d\n", reveal.Red, reveal.Green, reveal.Blue)
	}
	fmt.Println()
}
