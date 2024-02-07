package main

import (
	"reflect"
	"testing"
)

func TestGetTree(t *testing.T) {
	t.Run("Testing GetTree", func(t *testing.T) {
		start := RawNode{Key: "A", Left: "B", Right: "C"}
		nodeMapping := map[string]RawNode{
			"A": start,
			"B": {"B", "B", "B", nil},
			"C": {"C", "C", "C", nil},
		}
		expect := &Tree{
			Key: "A",
			Left: &Tree{
				Key: "B",
			},
			Right: &Tree{
				Key: "C",
			},
		}
		got := GetTree(start, nodeMapping)

		if !reflect.DeepEqual(expect, got) {
			t.Error("Trees don't match")

			t.Log("Got: ")
			PrettyPrint(got, "")

			t.Log("Expected: ")
			PrettyPrint(expect, "")

			t.Fail()
		}
	})

	t.Run("Testing GetTree", func(t *testing.T) {
		start := RawNode{Key: "A", Left: "B", Right: "C"}
		nodeMapping := map[string]RawNode{
			"A": start,
			"B": {"B", "D", "E", nil},
			"C": {"C", "Z", "G", nil},
			"D": {"D", "D", "D", nil},
			"E": {"E", "E", "E", nil},
			"Z": {"Z", "Z", "Z", nil},
			"G": {"G", "G", "G", nil},
		}
		expect := &Tree{
			Key: "A",
			Left: &Tree{
				Key: "B",
				Left: &Tree{
					Key: "D",
				},
				Right: &Tree{
					Key: "E",
				},
			},
			Right: &Tree{
				Key: "C",
				Left: &Tree{
					Key: "Z",
				},
				Right: &Tree{
					Key: "G",
				},
			},
		}
		got := GetTree(start, nodeMapping)

		if !reflect.DeepEqual(expect, got) {
			t.Error("Trees don't match")

			t.Log("Got: ")
			PrettyPrint(got, "")

			t.Log("Expected: ")
			PrettyPrint(expect, "")

			t.Fail()
		}
	})

	t.Run("Testing GetTree", func(t *testing.T) {
		start := RawNode{Key: "A", Left: "B", Right: "B"}
		nodeMapping := map[string]RawNode{
			"A": start,
			"B": {"B", "A", "Z", nil},
			"Z": {"Z", "Z", "Z", nil},
		}

		A := &Tree{
			Key: "A",
		}

		Z := &Tree{
			Key: "Z",
		}

		B := &Tree{
			Key:   "B",
			Left:  A,
			Right: Z,
		}

		A.Left = B
		A.Right = B

		expect := A
		got := GetTree(start, nodeMapping)
		// PrettyPrint(got, "")

		if !reflect.DeepEqual(expect, got) {
			t.Error("Trees don't match")

			t.Log("Got: ")
			// PrettyPrint(got, "")

			t.Log("Expected: ")
			// PrettyPrint(expect, "")

			t.Fail()
		}
	})
}
