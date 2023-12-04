package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"strconv"
	"strings"
)

func NewChallenge02() Challenge02 {
	return Challenge02{
		limits: map[string]int{
			"red":   12,
			"green": 13,
			"blue":  14,
		},
	}
}

type Challenge02 struct {
	limits map[string]int
}

func (c Challenge02) RunPartOne(input string) string {
	games := helpers.SplitLines(input)
	possibleGames := make([]int, 0)

	for _, game := range games {
		isGamePossible := true
		firstSplit := strings.Split(game, ":")
		grabs := strings.Split(firstSplit[1], ";")

		for _, grab := range grabs {
			grab := strings.TrimSpace(grab)
			cubes := strings.Split(grab, ",")
			for _, cube := range cubes {
				cube := strings.TrimSpace(cube)
				values := strings.Split(cube, " ")
				amount, _ := strconv.Atoi(values[0])
				color := values[1]

				if amount > c.limits[color] {
					isGamePossible = false
				}
			}
			if !isGamePossible {
				break
			}
		}

		if isGamePossible {
			gameData := strings.Split(firstSplit[0], " ")
			number, _ := strconv.Atoi(gameData[1])
			possibleGames = append(possibleGames, number)
		}
	}

	sum := aoc_math.Sum(possibleGames)

	return strconv.Itoa(sum)
}

func (c Challenge02) RunPartTwo(input string) string {
	games := helpers.SplitLines(input)
	powers := make([]int, 0)

	for _, game := range games {
		firstSplit := strings.Split(game, ":")
		grabs := strings.Split(firstSplit[1], ";")

		minimums := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		for _, grab := range grabs {
			grab := strings.TrimSpace(grab)
			cubes := strings.Split(grab, ",")
			for _, cube := range cubes {
				cube := strings.TrimSpace(cube)
				values := strings.Split(cube, " ")
				amount, _ := strconv.Atoi(values[0])
				color := values[1]

				if amount > minimums[color] {
					minimums[color] = amount
				}
			}
		}

		power := minimums["red"] * minimums["green"] * minimums["blue"]
		powers = append(powers, power)
	}

	return strconv.Itoa(aoc_math.Sum(powers))
}

func (c Challenge02) DataFolder() string {
	return "02"
}
