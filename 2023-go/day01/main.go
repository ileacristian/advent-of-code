package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

// Full problem text at: https://adventofcode.com/2023/day/1

func main() {
	inputFilename := "day01.txt"

	fmt.Println("First part: ", FirstPart(inputFilename))
	fmt.Println("Second part: ", SecondPart(inputFilename))

}

func FirstPart(filename string) int {
	return SumLines(filename, func(line string) int {
		return CalculateLine(line)
	})
}

func SecondPart(filename string) int {
	return SumLines(filename, func(line string) int {
		return CalculateLine(AdjustedLine(line))
	})
}

type LineCalculator func(line string) int

func SumLines(filename string, lineCalculatorFn LineCalculator) int {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		totalSum += lineCalculatorFn(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return totalSum
}

// Given a string xtwone3four it should become x2ne34
// The function finds digits expressed as words and replaces them with actual digits
// We are interested in only the start and the end of the string, so the function
// doesn't replace all words.
func AdjustedLine(line string) string {
	numbers := map[string]string{
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

	var keys []string
	for key := range numbers {
		keys = append(keys, key)
	}

	// replace only the first occurance of a word starting from the beginning
OuterLoop:
	for i := 0; i < len(line); i++ {
		// if we encounter a real digit, then we break
		// there's a sneaky bug if we don't break. Eg: xxx3xxxoneight -> xxx3xxx1ight (31) instead of xxx3xxxon8 (38)
		if line[i] >= '0' && line[i] <= '9' {
			break OuterLoop
		}
		for _, key := range keys {
			if strings.HasPrefix(line[i:], key) {
				line = strings.Replace(line, key, numbers[key], 1)
				break OuterLoop
			}
		}
	}

	// replace the first occurance of a word starting from the end
	for i := len(line) - 1; i >= 0; i-- {
		for _, key := range keys {
			if strings.HasPrefix(line[i:], key) {
				line = strings.Replace(line, key, numbers[key], 1)
				break
			}
		}
	}

	return line
}

// Finds the first digit and the last digit from a string and concatenates them into a new number
func CalculateLine(line string) int {
	runes := []rune(line)
	number := 0

	for _, character := range runes {
		if unicode.IsDigit(character) {
			digit := int(character - '0')
			number = number + (10 * digit)
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		if unicode.IsDigit(runes[i]) {
			digit := int(runes[i] - '0')
			number = number + digit
			break
		}
	}

	return number
}
