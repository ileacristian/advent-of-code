package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]rune
	rows := 0
	cols := 0

	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []rune(line))
		rows++
	}
	cols = len(matrix[0])

	fmt.Println(matrix, rows, cols)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(matrix, rows, cols))
	fmt.Printf("%v", matrix)
	// fmt.Println("Second Part: ", SecondPart(games))

}

func FirstPart(matrix [][]rune, rows, cols int) int {
	total := 0

	symbols := "#%&*+-/=@$"

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if strings.ContainsRune(symbols, matrix[i][j]) {
				total += LookAndSumParts(matrix, i, j)
			}
		}
	}

	return total
}

func LookAndSumParts(matrix [][]rune, row, col int) int {
	partsSum := 0

	partsSum += Check(matrix, row-1, col-1) // top left
	partsSum += Check(matrix, row-1, col)   // top
	partsSum += Check(matrix, row-1, col+1) // top right
	partsSum += Check(matrix, row, col+1)   // right
	partsSum += Check(matrix, row+1, col+1) // bottom right
	partsSum += Check(matrix, row+1, col)   // bottom
	partsSum += Check(matrix, row+1, col-1) // bottom left
	partsSum += Check(matrix, row, col-1)   // left

	return partsSum
}

func Check(matrix [][]rune, row, col int) int {
	if !validCoords(matrix, row, col) {
		return 0
	}

	value := matrix[row][col]
	if unicode.IsDigit(value) {
		return ExtractAndErase(matrix, row, col)
	}

	return 0
}

func ExtractAndErase(matrix [][]rune, row, col int) int {
	extractedDigits := string(matrix[row][col])
	matrix[row][col] = '.'

	// go left
	currentPosition := row - 1
	for currentPosition >= 0 {
		extractedDigits = string(matrix[currentPosition][col]) + extractedDigits
		matrix[currentPosition][col] = '.'
		currentPosition--
	}

	// go right
	currentPosition = row + 1
	for currentPosition < len(matrix[0]) {
		extractedDigits = extractedDigits + string(matrix[currentPosition][col])
		matrix[currentPosition][col] = '.'
		currentPosition++
	}

	extractedNumber, _ := strconv.Atoi(extractedDigits)
	return extractedNumber
}

func validCoords(matrix [][]rune, row, col int) bool {
	return row >= 0 && col >= 0 && row < len(matrix) && col < len(matrix[0])
}
