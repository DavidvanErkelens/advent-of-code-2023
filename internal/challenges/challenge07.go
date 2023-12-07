package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge07() Challenge07 {
	return Challenge07{
		cardScores: map[rune]int{
			'A': 13,
			'K': 12,
			'Q': 11,
			'J': 10,
			'T': 9,
			'9': 8,
			'8': 7,
			'7': 6,
			'6': 5,
			'5': 4,
			'4': 3,
			'3': 2,
			'2': 1,
		},
	}
}

type HandType int

const (
	FiveOfAKind  HandType = 6
	FourOfAKind  HandType = 5
	FullHouse    HandType = 4
	TwoPair      HandType = 2
	OnePair      HandType = 1
	ThreeOfAKind HandType = 3
	HighCard     HandType = 0
)

type Challenge07 struct {
	cardScores map[rune]int
}

type hand struct {
	cards []rune
	bid   int
}

func (h hand) getCounts() map[rune]int {
	counts := make(map[rune]int)
	for _, r := range h.cards {
		val := counts[r]
		counts[r] = val + 1
	}
	return counts
}

func (h hand) getType() HandType {
	counts := h.getCounts()
	pairs := 0
	threes := 0

	for _, v := range counts {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			threes += 1
		case 2:
			pairs += 1
		}
	}

	switch true {
	case pairs == 1 && threes == 1:
		return FullHouse
	case threes == 1:
		return ThreeOfAKind
	case pairs == 2:
		return TwoPair
	case pairs == 1:
		return OnePair
	}

	return HighCard
}

func (h hand) getTypePartTwo() HandType {
	counts := h.getCounts()
	pairs := 0
	threes := 0
	jokers := counts['J']

	for k, v := range counts {
		if k == 'J' {
			continue
		}

		if v == 5 {
			return FiveOfAKind
		}

		if v == 4 {
			if jokers == 1 {
				return FiveOfAKind
			}
			return FourOfAKind
		}

		if v == 3 {
			switch jokers {
			case 2:
				return FiveOfAKind
			case 1:
				return FourOfAKind
			}

			threes += 1
		}

		if v == 2 {
			switch jokers {
			case 3:
				return FiveOfAKind
			case 2:
				return FourOfAKind
			}

			pairs += 1
		}
	}

	switch true {
	case jokers >= 4:
		return FiveOfAKind

	case jokers == 3:
		return FourOfAKind

	case threes == 1 && pairs == 1,
		threes == 1 && jokers >= 1,
		pairs == 2 && jokers == 1,
		pairs == 1 && jokers >= 2:
		return FullHouse

	case pairs == 1 && jokers >= 1,
		jokers >= 2, threes == 1:
		return ThreeOfAKind

	case pairs == 2,
		pairs == 1 && jokers >= 1:
		return TwoPair

	case pairs >= 1,
		jokers == 1:
		return OnePair
	}

	return HighCard
}

func (h hand) getValue(mapping map[rune]int) int {
	return helpers.ReduceWithIndex(h.cards, func(index int, value rune, total int) int {
		return total + mapping[value]*aoc_math.IntPow(100, 4-index)
	}, 0)
}

func (c Challenge07) parseInput(input string) []hand {
	lines := helpers.SplitLines(input)
	hands := make([]hand, 0)

	for _, line := range lines {
		data := strings.Split(line, " ")
		hand := hand{
			cards: []rune(data[0]),
			bid:   helpers.NaiveStringToInt(data[1]),
		}

		hands = append(hands, hand)
	}

	return hands
}

func (c Challenge07) RunPartOne(input string) string {

	hands := c.parseInput(input)

	slices.SortFunc(hands, func(a, b hand) int {
		aType := a.getType()
		bType := b.getType()

		if aType != bType {
			return int(aType) - int(bType)
		}

		return a.getValue(c.cardScores) - b.getValue(c.cardScores)
	})

	total := 0

	for i, h := range hands {
		total += (i + 1) * h.bid
	}

	return strconv.Itoa(total)
}

func (c Challenge07) RunPartTwo(input string) string {
	hands := c.parseInput(input)

	valueMapping := c.cardScores
	valueMapping['J'] = 0

	slices.SortFunc(hands, func(a, b hand) int {
		aType := a.getTypePartTwo()
		bType := b.getTypePartTwo()

		if aType != bType {
			return int(aType) - int(bType)
		}

		return a.getValue(valueMapping) - b.getValue(valueMapping)
	})

	total := 0

	for i, h := range hands {
		total += (i + 1) * h.bid
	}

	return strconv.Itoa(total)
}

func (c Challenge07) DataFolder() string {
	return "07"
}
