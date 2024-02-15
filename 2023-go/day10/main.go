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
	VerticalPipe Tile = iota - 7
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

//go:embed day10.txt
var day10 string

func main() {
	// PrettyPrint(ParseInput(test1))
	// fmt.Println("")
	// PrettyPrint(ParseInput(test2))
	// fmt.Println("")
	// PrettyPrint(ParseInput(test3))
	// fmt.Println("")
	// PrettyPrint(ParseInput(test4))
	// fmt.Println("")
	// PrettyPrint(ParseInput(test5))
	pipes := ParseInput(day10)

	// PrettyPrint(pipes)

	fmt.Println()

	fmt.Println(FillAndGetMaxDistance(pipes))
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
	if t >= 0 {
		return fmt.Sprintf("%d", t)
	}

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

func FillAndGetMaxDistance(tiles [][]Tile) int {
	var start = FindStart(tiles)

	distances := make([][]Tile, len(tiles))
	for i := range tiles {
		distances[i] = make([]Tile, len(tiles[i]))
	}

	fmt.Println(start)

	var queue Queue[Pair[int]]
	var currentLocation = start
	distances[currentLocation.a][currentLocation.b] = 0
	queue.push(currentLocation)

	for !queue.empty() {
		currentLocation = queue.pop()
		newDirections := ValidDirectionsFrom(currentLocation, tiles)
		FillNewDirections(int(distances[currentLocation.a][currentLocation.b]+1), distances, newDirections)
		tiles[currentLocation.a][currentLocation.b] = Ground
		queue.push(newDirections...)
	}

	// PrettyPrint(distances)

	return int(distances[currentLocation.a][currentLocation.b])
}

func FindStart(tiles [][]Tile) (start Pair[int]) {
	for i := range tiles {
		for j := range tiles[0] {
			if tiles[i][j] == Start {
				start = Pair[int]{i, j}
			}
		}
	}
	return
}

func FillNewDirections(value int, tiles [][]Tile, newDirections []Pair[int]) {
	for _, newDir := range newDirections {
		tiles[newDir.a][newDir.b] = Tile(value)
	}
}

func ValidDirectionsFrom(loc Pair[int], tiles [][]Tile) []Pair[int] {
	line := loc.a
	col := loc.b
	output := []Pair[int]{}
	currTile := tiles[line][col]
	var left, right, up, down Tile = Ground, Ground, Ground, Ground

	if col-1 >= 0 {
		left = tiles[line][col-1]
	}
	if col+1 < len(tiles[0]) {
		right = tiles[line][col+1]
	}
	if line-1 >= 0 {
		up = tiles[line-1][col]
	}
	if line+1 < len(tiles) {
		down = tiles[line+1][col]
	}

	switch currTile {
	case Start:
		if up == VerticalPipe || up == TopLeftPipe || up == TopRightPipe {
			output = append(output, Pair[int]{line - 1, col})
		}
		if down == VerticalPipe || down == BottomLeftPipe || down == BottomRightPipe {
			output = append(output, Pair[int]{line + 1, col})
		}
		if left == HorizontalPipe || left == BottomLeftPipe || left == TopLeftPipe {
			output = append(output, Pair[int]{line, col - 1})
		}
		if right == HorizontalPipe || right == BottomRightPipe || right == TopRightPipe {
			output = append(output, Pair[int]{line, col + 1})
		}
	case VerticalPipe:
		if up == VerticalPipe || up == TopLeftPipe || up == TopRightPipe {
			output = append(output, Pair[int]{line - 1, col})
		}
		if down == VerticalPipe || down == BottomLeftPipe || down == BottomRightPipe {
			output = append(output, Pair[int]{line + 1, col})
		}
	case HorizontalPipe:
		if left == HorizontalPipe || left == BottomLeftPipe || left == TopLeftPipe {
			output = append(output, Pair[int]{line, col - 1})
		}
		if right == HorizontalPipe || right == BottomRightPipe || right == TopRightPipe {
			output = append(output, Pair[int]{line, col + 1})
		}
	case TopLeftPipe:
		if down == VerticalPipe || down == BottomRightPipe || down == BottomLeftPipe {
			output = append(output, Pair[int]{line + 1, col})
		}
		if right == HorizontalPipe || right == BottomRightPipe || right == TopRightPipe {
			output = append(output, Pair[int]{line, col + 1})
		}
	case TopRightPipe:
		if left == HorizontalPipe || left == BottomLeftPipe || left == TopLeftPipe {
			output = append(output, Pair[int]{line, col - 1})
		}
		if down == VerticalPipe || down == BottomRightPipe || down == BottomLeftPipe {
			output = append(output, Pair[int]{line + 1, col})
		}
	case BottomLeftPipe:
		if up == VerticalPipe || up == TopLeftPipe || up == TopRightPipe {
			output = append(output, Pair[int]{line - 1, col})
		}
		if right == HorizontalPipe || right == BottomRightPipe || right == TopRightPipe {
			output = append(output, Pair[int]{line, col + 1})
		}
	case BottomRightPipe:
		if left == HorizontalPipe || left == BottomLeftPipe || left == TopLeftPipe {
			output = append(output, Pair[int]{line, col - 1})
		}
		if up == VerticalPipe || up == TopLeftPipe || up == TopRightPipe {
			output = append(output, Pair[int]{line - 1, col})
		}
	}

	return output
}

type Pair[T any] struct {
	a, b T
}

type Stack[Elem any] struct {
	_stack []Elem
}

func (s *Stack[Elem]) push(e ...Elem) {
	s._stack = append(s._stack, e...)
}

func (s *Stack[Elem]) pop() Elem {
	popped := s._stack[len(s._stack)-1]
	s._stack = s._stack[:len(s._stack)-1]
	return popped
}

func (s Stack[Elem]) empty() bool {
	return len(s._stack) == 0
}

type Queue[Elem any] struct {
	_queue []Elem
}

func (s *Queue[Elem]) push(e ...Elem) {
	s._queue = append(s._queue, e...)
}

func (q *Queue[Elem]) pop() Elem {
	popped := q._queue[0]
	q._queue = q._queue[1:]
	return popped
}

func (q Queue[Elem]) empty() bool {
	return len(q._queue) == 0
}
