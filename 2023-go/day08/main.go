package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Tree struct {
	Key         string
	Left, Right *Tree
	Visited     bool
}

func PrettyPrint(node *Tree, prefix string) {
	if node == nil {
		return
	}

	// Print the current node
	fmt.Println(prefix + node.Key)

	if node.Visited {
		return
	}

	node.Visited = true

	// Prepare the prefix for the child nodes
	newPrefix := prefix + "|\t"

	// Recursively print left and right subtree
	PrettyPrint(node.Left, newPrefix)
	PrettyPrint(node.Right, newPrefix)
}

type Direction int

const (
	Left Direction = iota
	Right
)

type RawNode struct {
	Key   string
	Left  string
	Right string
	Tree  *Tree
}

func main() {
	file, err := os.Open("day08.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var directions []Direction
	nodes := []RawNode{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if len(directions) == 0 {
			directions = ParseDirectionsLine(line)
			continue
		}

		nodes = append(nodes, ParseNodeLine(line))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	tree := TreeFromNodeWithKey(nodes, "AAA")
	fmt.Println("First Part: ", FirstPart(directions, tree))

	trees := TreesWithASuffix(nodes)
	fmt.Println("Second Part: ", SecondPart(directions, trees))
}

func ParseDirectionsLine(line string) []Direction {
	directions := make([]Direction, len(line))
	for i, d := range line {
		switch d {
		case 'L':
			directions[i] = Left
		case 'R':
			directions[i] = Right
		default:
			panic("Unknown direction")
		}
	}

	return directions
}

var lineRegexp = regexp.MustCompile(`([A-Z]{3}) = \(([A-Z]{3}), ([A-Z]{3})\)`)

func ParseNodeLine(line string) RawNode {

	matches := lineRegexp.FindStringSubmatch(line)

	if len(matches) != 4 {
		panic("Regexpt Parsing issue")
	}

	return RawNode{Key: matches[1], Left: matches[2], Right: matches[3]}
}

func TreeFromNodeWithKey(nodes []RawNode, key string) *Tree {
	nodeMapping := make(map[string]RawNode, len(nodes))
	for _, node := range nodes {
		nodeMapping[node.Key] = node
	}

	startNode := nodeMapping[key]
	tree := GetTree(startNode, nodeMapping)
	return tree
}

func GetTree(node RawNode, nodeMapping map[string]RawNode) *Tree {
	left := nodeMapping[node.Left]
	right := nodeMapping[node.Right]

	if node.Tree != nil {
		return node.Tree
	}
	var newTree *Tree = &Tree{Key: node.Key}
	node.Tree = newTree
	nodeMapping[node.Key] = node

	if node.Key == left.Key && node.Key == right.Key {
		return newTree
	}

	if node.Key == left.Key {
		newTree.Right = GetTree(right, nodeMapping)
		return newTree
	}

	if node.Key == right.Key {
		newTree.Left = GetTree(left, nodeMapping)
		return newTree
	}

	newTree.Left = GetTree(left, nodeMapping)
	newTree.Right = GetTree(right, nodeMapping)
	return newTree
}

func FirstPart(directions []Direction, tree *Tree) int {
	distance := 0

	currentDirection := 0
	for tree.Key != "ZZZ" {
		if directions[currentDirection] == Left {
			tree = tree.Left
		} else {
			tree = tree.Right
		}

		if tree == nil {
			panic("Should not get here")
		}

		currentDirection = (currentDirection + 1) % len(directions)
		distance++
	}

	return distance
}

func TreesWithASuffix(nodes []RawNode) []*Tree {
	trees := []*Tree{}
	for _, node := range nodes {
		if strings.HasSuffix(node.Key, "A") {
			trees = append(trees, TreeFromNodeWithKey(nodes, node.Key))
		}
	}
	return trees
}

func SecondPart(directions []Direction, trees []*Tree) int {
	partialResults := make([]int, len(trees))
	var wg sync.WaitGroup

	for i, tree := range trees {
		wg.Add(1)
		go DistanceToZ(directions, tree, &partialResults[i], &wg)
	}

	wg.Wait()

	result := lcm(partialResults[0], partialResults[1])

	for i := 2; i < len(partialResults); i++ {
		result = lcm(result, partialResults[i])
	}

	return result
}

func DistanceToZ(directions []Direction, tree *Tree, result *int, wg *sync.WaitGroup) {
	distance := 0

	currentDirection := 0
	for !strings.HasSuffix(tree.Key, "Z") {
		if directions[currentDirection] == Left {
			tree = tree.Left
		} else {
			tree = tree.Right
		}

		if tree == nil {
			panic("Should not get here")
		}

		currentDirection = (currentDirection + 1) % len(directions)
		distance++
	}
	*result = distance
	wg.Done()
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
