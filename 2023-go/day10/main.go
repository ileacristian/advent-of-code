package main

import (
	_ "embed"
	"fmt"
	"strings"
)

const (
	TopLeftCornerC     = "╔"
	TopRightCornerC    = "╗"
	BottomLeftCornerC  = "╚"
	BottomRightCornerC = "╝"
	VerticalC          = "║"
	HorizontalC        = "═"
	GroundC            = "░"
	StartC             = "█"
)

type Tile int

const (
	VerticalPipe = iota
	HorizontalPipe
	BottomLeftPipe
	BottomRightPipe
	TopRightPipe
	TopLeftPipe
	Ground
	Start
)

//go:embed test1.txt
var test1 string

//go:embed test2.txt
var test2 string

//go:embed test3.txt
var test3 string

//go:embed test4.txt
var test4 string

//go:embed test5.txt
var test5 string

func main() {
	PrettyPrint(ParseInput(test1))
	fmt.Println("")
	PrettyPrint(ParseInput(test2))
	fmt.Println("")
	PrettyPrint(ParseInput(test3))
	fmt.Println("")
	PrettyPrint(ParseInput(test4))
	fmt.Println("")
	PrettyPrint(ParseInput(test5))
}

func ParseInput(surfaceMapString string) [][]Tile {
	lines := strings.Split(surfaceMapString, "\n")
	output := make([][]Tile, len(lines))

	for i, line := range lines {
		output[i] = make([]Tile, len(line))
		for j, r := range line {
			output[i][j] = runeToTile(r)
		}
	}

	return output
}

func runeToTile(r rune) Tile {
	runeTileMap := map[rune]Tile{
		'|': VerticalPipe,
		'-': HorizontalPipe,
		'L': BottomLeftPipe,
		'J': BottomRightPipe,
		'7': TopRightPipe,
		'F': TopLeftPipe,
		'.': Ground,
		'S': Start,
	}
	return runeTileMap[r]
}

func TileToPrettyChar(t Tile) string {
	tileCharMap := map[Tile]string{
		VerticalPipe:    VerticalC,
		HorizontalPipe:  HorizontalC,
		TopLeftPipe:     TopLeftCornerC,
		TopRightPipe:    TopRightCornerC,
		BottomLeftPipe:  BottomLeftCornerC,
		BottomRightPipe: BottomRightCornerC,
		Start:           StartC,
		Ground:          GroundC,
	}

	return tileCharMap[t]
}

// example outputs:
// 1)
// ═╚║╔╗
// ╗█═╗║
// ╚║╗║║
// ═╚═╝║
// ╚║═╝╔
//
// 2)
// ░░╔╗░
// ░╔╝║░
// █╝░╚╗
// ║╔══╝
// ╚╝░░░
func PrettyPrint(tiles [][]Tile) {
	for _, line := range tiles {
		for _, tile := range line {
			fmt.Printf("%s", TileToPrettyChar(tile))
		}
		fmt.Println()
	}
}
