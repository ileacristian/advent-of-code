package main

import (
	"bufio"
	"fmt"
	"os"
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
	// fmt.Println("Second Part: ", SecondPart(games))

}

func FirstPart(matrix [][]rune, rows, cols int) int {
	total := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 5 {
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
	value := matrix[row][col]
	if unicode.IsDigit(value) {
		return ExtractAndErase(matrix, row, col)
	}

	return 0
}

func ExtractAndErase(matrix [][]rune, row, col int) int {
	// TODO
	return 0
}
