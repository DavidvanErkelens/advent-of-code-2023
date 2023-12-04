package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"advent-of-code-2023/internal/helpers/aoc_range"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge04() Challenge04 {
	return Challenge04{}
}

type Challenge04 struct {
}

func (c Challenge04) RunPartOne(input string) string {
	lines := helpers.SplitLines(input)
	scoreSums := 0

	for _, line := range lines {
		configDataSplit := strings.Split(line, ":")
		numbersSplit := strings.Split(configDataSplit[1], "|")

		cardNumbers := helpers.StringListOfNumericValuesToSlice(numbersSplit[0], " ")
		winningNumbers := helpers.StringListOfNumericValuesToSlice(numbersSplit[1], " ")

		winners := helpers.Reduce(cardNumbers, func(cardNumber int, total int) int {
			return total + helpers.BoolToInt(slices.Contains(winningNumbers, cardNumber))
		}, 0)

		if winners > 0 {
			outcome := aoc_math.IntPow(2, winners-1)
			scoreSums += outcome
		}
	}

	return strconv.Itoa(scoreSums)
}

func (c Challenge04) RunPartTwo(input string) string {
	lines := helpers.SplitLines(input)
	numberOfCards := make(map[int]int, len(lines)+1)

	for _, idx := range aoc_range.NewRange(1, len(lines)) {
		numberOfCards[idx] = 1
	}

	for idx, line := range lines {
		configDataSplit := strings.Split(line, ":")
		numbersSplit := strings.Split(configDataSplit[1], "|")

		cardNumbers := helpers.StringListOfNumericValuesToSlice(numbersSplit[0], " ")
		winningNumbers := helpers.StringListOfNumericValuesToSlice(numbersSplit[1], " ")

		winners := helpers.Reduce(cardNumbers, func(cardNumber int, total int) int {
			return total + helpers.BoolToInt(slices.Contains(winningNumbers, cardNumber))
		}, 0)

		if winners == 0 {
			continue
		}

		cardNumber := idx + 1
		cardsForThisNumber := numberOfCards[cardNumber]

		cardsToAdd := aoc_range.NewRange(cardNumber+1, cardNumber+winners)

		for _, extraCardIdx := range cardsToAdd {
			numberOfCards[extraCardIdx] += cardsForThisNumber
		}
	}

	totalCards := 0
	for _, cards := range numberOfCards {
		totalCards += cards
	}

	return strconv.Itoa(totalCards)
}

func (c Challenge04) DataFolder() string {
	return "04"
}
