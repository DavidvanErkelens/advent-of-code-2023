package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"strconv"
	"strings"
)

func NewChallenge08() Challenge08 {
	return Challenge08{}
}

type Challenge08 struct {
}

func (c Challenge08) parseInput(input string) ([]rune, map[string]map[rune]string) {
	ps := helpers.SplitParagraphs(input)
	instructions := []rune(ps[0])

	lines := helpers.SplitLines(ps[1])
	mapping := make(map[string]map[rune]string)

	for _, line := range lines {
		fromTo := strings.Split(line, " = ")
		from := fromTo[0]
		leftRight := strings.Split(fromTo[1], ", ")
		left := strings.TrimLeft(leftRight[0], "(")
		right := strings.TrimRight(leftRight[1], ")")

		lineMap := map[rune]string{
			'L': left,
			'R': right,
		}

		mapping[from] = lineMap
	}

	return instructions, mapping
}

func (c Challenge08) RunPartOne(input string) string {
	instructions, mapping := c.parseInput(input)

	current := "AAA"
	steps := 0

	for current != "ZZZ" {
		instruction := instructions[steps%len(instructions)]
		steps += 1
		current = mapping[current][instruction]
	}

	return strconv.Itoa(steps)
}

func (c Challenge08) RunPartTwo(input string) string {
	instructions, mapping := c.parseInput(input)

	currentNodes := make([]string, 0)

	for key, _ := range mapping {
		if strings.HasSuffix(key, "A") {
			currentNodes = append(currentNodes, key)
		}
	}

	steps := 0
	zReached := make([]int, 0)

	for len(currentNodes) > 0 {
		instruction := instructions[steps%len(instructions)]
		steps += 1

		currentNodes = helpers.Map(currentNodes, func(node string) string {
			return mapping[node][instruction]
		})

		nodesLeft := make([]string, 0)

		for _, node := range currentNodes {
			if strings.HasSuffix(node, "Z") {
				zReached = append(zReached, steps)
			} else {
				nodesLeft = append(nodesLeft, node)
			}
		}

		currentNodes = nodesLeft
	}

	return strconv.Itoa(aoc_math.LCM(zReached...))
}

func (c Challenge08) DataFolder() string {
	return "08"
}
