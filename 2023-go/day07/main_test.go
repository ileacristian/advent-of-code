package main

import "testing"

func TestDetermineHandType(t *testing.T) {
	t.Run("Test determineHandType to determine correct hand type when jokers are not wildcards", func(t *testing.T) {
		hasJoker := false

		cardAndExpected := map[string]HandType{
			"32T3K": OnePair,
			"T55J5": ThreeOfAKind,
			"KK677": TwoPair,
			"KTJJT": TwoPair,
			"QQQJA": ThreeOfAKind,
			"QQJQQ": FourOfAKind,
			"AAAAA": FiveOfAKind,
			"23456": HighCard,
			"AJ4A4": TwoPair,
			"3QTJJ": OnePair,
		}

		for cards, expectedHandType := range cardAndExpected {
			hand := NewHand(cards, 0)
			hand.determineHandType(hasJoker)
			got := hand.Type
			if got != expectedHandType {
				t.Errorf("Expected handType %v for card: %s - got %v instead", expectedHandType, cards, got)
			}
		}

	})

	t.Run("Test NewHandType to determine correct hand type when jokers are not wildcards", func(t *testing.T) {
		hasJoker := true

		cardAndExpected := map[string]HandType{
			"32T3K": OnePair,
			"T55J5": FourOfAKind,
			"KK677": TwoPair,
			"KTJJT": FourOfAKind,
			"QQQJA": FourOfAKind,
			"QQJQQ": FiveOfAKind,
			"AAAAA": FiveOfAKind,
			"23456": HighCard,
			"JTJJT": FiveOfAKind,
			"AJ4A4": FullHouse,
			"3QTJJ": ThreeOfAKind,
		}

		for cards, expectedHandType := range cardAndExpected {
			if cards == "3QTJJ" {
				print("caca")
			}
			hand := NewHand(cards, 0)
			hand.determineHandType(hasJoker)
			got := hand.Type
			if got != expectedHandType {
				t.Errorf("Expected handType %v for card: %s - got %v instead", expectedHandType, cards, got)
			}
		}

	})
}
