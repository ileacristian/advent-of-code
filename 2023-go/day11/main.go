package main

import (
	_ "embed"
	"fmt"
	"slices"
	"strings"
)

type Tile int

const (
	Space Tile = iota
	Galaxy
)

//go:embed test.txt
var test string

func main() {
	universe := parseInput(test)
	prettyPrintUniverse(universe)
	expandUniverse(&universe)
	fmt.Println()
	prettyPrintUniverse(universe)
}

func parseInput(input string) [][]Tile {
	lines := strings.Split(input, "\n")

	output := make([][]Tile, len(lines))

	for i, line := range lines {
		output[i] = make([]Tile, len(line))
		for j, r := range line {
			if r == '.' {
				output[i][j] = Space

			} else {
				output[i][j] = Galaxy
			}
		}
	}

	return output
}

func expandUniverse(tiles *[][]Tile) {
	// lines
	emptyLinesIndices := []int{}
	for i, line := range *tiles {
		emptySpace := true
		for _, tile := range line {
			if tile == Galaxy {
				emptySpace = false
			}
		}
		if emptySpace {
			emptyLinesIndices = append(emptyLinesIndices, i)
		}
	}

	offset := 0
	for _, emptyLineIdx := range emptyLinesIndices {
		insertRowAt(tiles, emptyLineIdx+offset)
		offset++
	}

	// columns
	emptyColumns := []int{}
	for col := range (*tiles)[0] {
		emptySpace := true
		for i := range *tiles {
			if (*tiles)[i][col] == Galaxy {
				emptySpace = false
			}
		}
		if emptySpace {
			emptyColumns = append(emptyColumns, col)
		}
	}

	offset = 0
	for _, emptyColumn := range emptyColumns {
		insertColAt(tiles, emptyColumn+offset)
		offset++
	}
}

func insertRowAt(tiles *[][]Tile, n int) {
	newRow := make([]Tile, len((*tiles)[0]))
	for i := range newRow {
		newRow[i] = Space
	}

	*tiles = slices.Insert(*tiles, n, newRow)
}

func insertColAt(tiles *[][]Tile, n int) {
	for i := range *tiles {
		(*tiles)[i] = slices.Insert((*tiles)[i], n, Space)
	}
}

func prettyPrintUniverse(universe [][]Tile) {
	for i := range universe {
		for j := range universe[i] {
			switch universe[i][j] {
			case Galaxy:
				fmt.Print("#")
			case Space:
				fmt.Print(".")
			default:
				fmt.Print(universe[i][j])
			}
		}
		fmt.Println()
	}
}
