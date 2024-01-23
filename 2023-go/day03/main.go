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

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	print(matrix)
	copiedMatrix := make([][]rune, len(matrix))
	copy(copiedMatrix, matrix)
	// fmt.Println("First Part: ", FirstPart(matrix, rows, cols))
	fmt.Println("Second Part: ", SecondPart(copiedMatrix, rows, cols))

	// print(matrix)

}

func FirstPart(matrix [][]rune, rows, cols int) int {
	total := 0

	symbols := "#%&*+-/=@$"

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if strings.ContainsRune(symbols, matrix[i][j]) {
				matrix[i][j] = '.'
				total += LookAndSumParts(matrix, i, j)
			}
		}
	}

	return total
}

func SecondPart(matrix [][]rune, rows, cols int) int {
	total := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '*' && hasOnlyTwoNeighbors(matrix, i, j) {
				matrix[i][j] = '.'
				total += lookAndMulParts(matrix, i, j)
			}
		}
	}
	return total
}

func around(row, col int) [][]int {
	return [][]int{
		{row - 1, col - 1}, // top left
		{row - 1, col},     // top
		{row - 1, col + 1}, // top right
		{row, col + 1},     // right
		{row + 1, col + 1}, // bottom right
		{row + 1, col},     // bottom
		{row + 1, col - 1}, // bottom left
		{row, col - 1},     // left
	}
}

func hasOnlyTwoNeighbors(matrix [][]rune, row, col int) bool {
	neighbors := 0
	for _, pair := range around(row, col) {
		if unicode.IsDigit(matrix[pair[0]][pair[1]]) {
			neighbors++
			fmt.Println(row, col, neighbors)
		}
	}

	return neighbors == 2
}

func lookAndMulParts(matrix [][]rune, row, col int) int {
	partsProd := 1

	for _, pair := range around(row, col) {
		partsProd *= Check(matrix, pair[0], pair[1])
	}

	return partsProd
}

func LookAndSumParts(matrix [][]rune, row, col int) int {
	partsSum := 0

	for _, pair := range around(row, col) {
		partsSum += Check(matrix, pair[0], pair[1])
	}

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
	currentPosition := col - 1
	for currentPosition >= 0 && unicode.IsDigit(matrix[row][currentPosition]) {
		extractedDigits = string(matrix[row][currentPosition]) + extractedDigits
		matrix[row][currentPosition] = '.'
		currentPosition--
	}

	// go right
	currentPosition = col + 1
	for currentPosition < len(matrix[0]) && unicode.IsDigit(matrix[row][currentPosition]) {
		extractedDigits = extractedDigits + string(matrix[row][currentPosition])
		matrix[row][currentPosition] = '.'
		currentPosition++
	}

	extractedNumber, _ := strconv.Atoi(extractedDigits)
	return extractedNumber
}

func validCoords(matrix [][]rune, row, col int) bool {
	return row >= 0 && col >= 0 && row < len(matrix) && col < len(matrix[0])
}

func print(matrix [][]rune) {
	for _, row := range matrix {
		fmt.Println(string(row))
	}
}
