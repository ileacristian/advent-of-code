package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	WinningNumbers   []int
	CandidateNumbers []int
}

func main() {
	file, err := os.Open("day04.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cards := []ScratchCard{}
	for scanner.Scan() {
		line := scanner.Text()
		cards = append(cards, parseLine(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(cards))
	fmt.Println("Second Part: ", SecondPart(cards))
}

func FirstPart(cards []ScratchCard) int {
	totalPoints := 0
	for _, card := range cards {
		points := 0
		for _, number := range card.CandidateNumbers {
			if contains(card.WinningNumbers, number) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}

	return totalPoints
}

func SecondPart(cards []ScratchCard) int {
	cardsCounts := make([]int, len(cards))

	for i := range cardsCounts {
		cardsCounts[i] = 1
	}

	for i, card := range cards {
		extraCards := matches(card)
		for j := i + 1; j < i+1+extraCards; j++ {
			cardsCounts[j] += 1 * cardsCounts[i]
		}
	}

	sum := 0
	for _, value := range cardsCounts {
		sum += value
	}
	return sum
}

func matches(card ScratchCard) int {
	matches := 0
	for _, number := range card.CandidateNumbers {
		if contains(card.WinningNumbers, number) {
			matches++
		}
	}

	return matches
}

func parseLine(line string) ScratchCard {
	re := regexp.MustCompile(`Card\s+\d+: ([0-9\s]*) \| ([0-9\s]*)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) != 3 {
		fmt.Println(line)
		fmt.Println(matches)
		panic("Cannot parse input")
	}

	winningNumbersStr := strings.Fields(matches[1])
	candidateNumbersStr := strings.Fields(matches[2])

	winningNumbers := intsFromStrs(winningNumbersStr)
	candidateNumbers := intsFromStrs(candidateNumbersStr)

	return ScratchCard{
		WinningNumbers:   winningNumbers,
		CandidateNumbers: candidateNumbers,
	}
}

func intsFromStrs(strNumbers []string) []int {
	ints := make([]int, len(strNumbers))
	for i, s := range strNumbers {
		if num, err := strconv.Atoi(s); err == nil {
			ints[i] = num
		}
	}

	return ints
}

func contains(slice []int, number int) bool {
	for _, v := range slice {
		if v == number {
			return true
		}
	}
	return false
}
