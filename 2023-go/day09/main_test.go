package main

import (
	"fmt"
	"testing"
)

func TestNextElement(t *testing.T) {
	tests := map[int][]int{
		18: {0, 3, 6, 9, 12, 15},
		28: {1, 3, 6, 10, 15, 21},
		68: {10, 13, 16, 21, 30, 45},
	}

	for expected, input := range tests {
		testName := fmt.Sprintf("Test with input: %v and expected: %d", input, expected)
		t.Run(testName, func(t *testing.T) {
			got := nextElement(input)
			if expected != got {
				t.Errorf("Got %d but expected %d", got, expected)
			}
		})
	}

}

func TestPrevElement(t *testing.T) {
	tests := map[int][]int{
		-3: {0, 3, 6, 9, 12, 15},
		0:  {1, 3, 6, 10, 15, 21},
		5:  {10, 13, 16, 21, 30, 45},
	}

	for expected, input := range tests {
		testName := fmt.Sprintf("Test with input: %v and expected: %d", input, expected)
		t.Run(testName, func(t *testing.T) {
			got := prevElement(input)
			if expected != got {
				t.Errorf("Got %d but expected %d", got, expected)
			}
		})
	}

}
