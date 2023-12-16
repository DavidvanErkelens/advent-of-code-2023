package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"strconv"
	"strings"
)

func NewChallenge15() Challenge15 {
	return Challenge15{}
}

type Challenge15 struct {
}

func (c Challenge15) calculateHash(elem string) int {
	score := 0

	for _, r := range elem {
		score += int(r)
		score *= 17
		score %= 256
	}

	return score
}

func (c Challenge15) RunPartOne(input string) string {
	totalSum := 0
	parts := strings.Split(input, ",")
	for _, part := range parts {
		totalSum += c.calculateHash(part)
	}
	return strconv.Itoa(totalSum)
}

type box struct {
	lenses []lens
}

type lens struct {
	label string
	value int
}

func (c Challenge15) RunPartTwo(input string) string {
	parts := strings.Split(input, ",")
	boxes := make([]box, 256)

	for _, part := range parts {
		if strings.Contains(part, "=") {
			data := strings.Split(part, "=")
			value := helpers.NaiveStringToInt(data[1])
			box := c.calculateHash(data[0])
			found := false

			for i := 0; i < len(boxes[box].lenses); i++ {
				if boxes[box].lenses[i].label == data[0] {
					boxes[box].lenses[i].value = value
					found = true
				}

			}

			if !found {
				boxes[box].lenses = append(boxes[box].lenses, lens{label: data[0], value: value})
			}

		} else {
			toRemove := strings.TrimRight(part, "-")
			for boxIdx := 0; boxIdx < len(boxes); boxIdx++ {
				for i, lens := range boxes[boxIdx].lenses {
					if lens.label == toRemove {
						boxes[boxIdx].lenses = helpers.RemoveIndex(boxes[boxIdx].lenses, i)
						break
					}
				}
			}
		}
	}

	score := 0

	for boxIdx, box := range boxes {
		for lensIdx, lens := range box.lenses {
			lScore := (boxIdx + 1) * (lensIdx + 1) * lens.value
			score += lScore
		}
	}

	return strconv.Itoa(score)
}

func (c Challenge15) DataFolder() string {
	return "15"
}
