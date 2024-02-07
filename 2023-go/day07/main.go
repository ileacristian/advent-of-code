package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type KeyValuePair struct {
	Key   rune
	Value int
}

func Pairs(m map[rune]int) []KeyValuePair {
	pairs := make([]KeyValuePair, 0, len(m))
	for key, value := range m {
		pairs = append(pairs, KeyValuePair{Key: key, Value: value})
	}
	return pairs
}

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func (t HandType) String() string {
	return [...]string{"HighCard", "OnePair", "TwoPair", "ThreeOfAKind", "FullHouse", "FourOfAKind", "FiveOfAKind"}[t-1]
}

func NewHandType(cards string, hasJoker bool) HandType {
	if len(cards) != 5 {
		panic("Unexpected number of cards")
	}

	cardCount := make(map[rune]int)
	for _, card := range cards {
		cardCount[card]++
	}

	// ignore jokers for initial hand type
	if _, hasJ := cardCount['J']; hasJoker && hasJ {
		cardCount['J'] = 0
	}

	cardCountKVP := Pairs(cardCount)

	switch len(cardCountKVP) {
	case 1:
		return FiveOfAKind
	case 2:
		if cardCountKVP[0].Value == 4 || cardCountKVP[1].Value == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	case 3:
		if cardCountKVP[0].Value == 3 || cardCountKVP[1].Value == 3 || cardCountKVP[2].Value == 3 {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	case 4:
		if cardCountKVP[0].Value < 2 && cardCountKVP[1].Value < 2 && cardCountKVP[2].Value < 2 && cardCountKVP[3].Value < 2 {
			return HighCard
		} else {
			return OnePair
		}
	case 5:
		return HighCard
	}

	return HighCard
}

type Hand struct {
	Cards string
	Bid   int
	Type  HandType
}

func NewHand(cards string, bid int) Hand {
	hand := Hand{Cards: cards, Bid: bid}
	return hand
}

func (h Hand) LessThan(other Hand, hasJoker bool) bool {
	if h.Type == other.Type {
		return h.SortableCards(hasJoker) < other.SortableCards(hasJoker)
	} else {
		return h.Type < other.Type
	}
}

func (h Hand) SortableCards(hasJoker bool) string {
	// Remapping cards to ascii characters so they sort using "<" according to their game value
	mapping := map[rune]rune{
		'A': 'Z',
		'K': 'Y',
		'Q': 'X',
		'J': 'W',
		'T': 'V',
	}

	if hasJoker {
		mapping['J'] = '0'
	}

	cards := h.Cards
	for k, v := range mapping {
		cards = strings.ReplaceAll(cards, string(k), string(v))
	}

	return cards
}

func (h *Hand) determineHandType(hasJoker bool) {
	h.Type = NewHandType(h.Cards, hasJoker)
	if hasJoker {
		h.adjustRankForJokers()
	}
}

func (h *Hand) adjustRankForJokers() {
	jokerCount := countOccurrences(h.Cards, 'J')

	for jokerCount > 0 {
		switch h.Type {
		case HighCard:
			h.Type = OnePair
		case OnePair:
			h.Type = ThreeOfAKind
		case TwoPair:
			h.Type = FullHouse
		case FullHouse:
			h.Type = FourOfAKind
		case ThreeOfAKind:
			h.Type = FourOfAKind
		case FourOfAKind:
			h.Type = FiveOfAKind
		}
		jokerCount--
	}
}

func countOccurrences(s string, char rune) int {
	count := 0
	for _, c := range s {
		if c == char {
			count++
		}
	}
	return count
}

func main() {
	file, err := os.Open("day07.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, " ")
		cards := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, NewHand(cards, bid))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("First Part: ", FirstPart(hands))
	fmt.Println("Second Part: ", SecondPart(hands))
}

type ByRankNoJoker []Hand

func (h ByRankNoJoker) Len() int           { return len(h) }
func (h ByRankNoJoker) Less(i, j int) bool { return h[i].LessThan(h[j], false) }
func (h ByRankNoJoker) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type ByRankWithJoker []Hand

func (h ByRankWithJoker) Len() int           { return len(h) }
func (h ByRankWithJoker) Less(i, j int) bool { return h[i].LessThan(h[j], true) }
func (h ByRankWithJoker) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func FirstPart(hands []Hand) int {
	total := 0
	hasJoker := false
	for i := 0; i < len(hands); i++ {
		hands[i].determineHandType(hasJoker)
	}
	sort.Sort(ByRankNoJoker(hands))
	for i, h := range hands {
		total += (i + 1) * h.Bid
	}
	return total
}

func SecondPart(hands []Hand) int {
	total := 0
	hasJoker := true

	for i := 0; i < len(hands); i++ {
		hands[i].determineHandType(hasJoker)
	}

	sort.Sort(ByRankWithJoker(hands))
	for i, h := range hands {
		total += (i + 1) * h.Bid
	}
	return total
}
