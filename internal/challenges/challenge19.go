package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_range"
	"maps"
	"strconv"
	"strings"
)

func NewChallenge19() Challenge19 {
	return Challenge19{
		register: make(map[string]int),
	}
}

type Challenge19 struct {
	register map[string]int
}

type workflow struct {
	rules    []rule
	fallback string
}

type rule struct {
	variable    string
	operator    func(int, int) bool
	operatorStr string
	value       int
	destination string
}

type rangeContainer struct {
	state  string
	ranges map[string]aoc_range.Range
}

func (c Challenge19) isTrue(r rule) bool {
	return r.operator(c.register[r.variable], r.value)
}

func (c Challenge19) parseInput(input string) map[string]workflow {
	lines := helpers.SplitLines(input)
	workflows := make(map[string]workflow)
	smallerThanFn := func(input, compare int) bool { return input < compare }
	largerThanFn := func(input, compare int) bool { return input > compare }

	for _, wf := range lines {
		nameDataSplit := strings.Split(wf, "{")
		name := nameDataSplit[0]
		rules := strings.Split(strings.TrimRight(nameDataSplit[1], "}"), ",")
		parsedRules := make([]rule, len(rules)-1)
		var finalDestination string

		for i, r := range rules {
			formulaDestination := strings.Split(r, ":")

			if len(formulaDestination) == 1 {
				finalDestination = r
				continue
			}

			formula := formulaDestination[0]
			destination := formulaDestination[1]

			if strings.Index(formula, ">") > 0 {
				split := strings.Split(formula, ">")
				parsedRules[i] = rule{
					variable:    split[0],
					operator:    largerThanFn,
					operatorStr: ">",
					value:       helpers.NaiveStringToInt(split[1]),
					destination: destination,
				}
				continue
			}

			if strings.Index(formula, "<") > 0 {
				split := strings.Split(formula, "<")
				parsedRules[i] = rule{
					variable:    split[0],
					operator:    smallerThanFn,
					operatorStr: "<",
					value:       helpers.NaiveStringToInt(split[1]),
					destination: destination,
				}
				continue
			}

			panic("this should not be reachable")
		}

		workflows[name] = workflow{
			rules:    parsedRules,
			fallback: finalDestination,
		}
	}

	return workflows
}

func (c Challenge19) setRegister(line string) {
	line = strings.Trim(line, "{}")
	vars := strings.Split(line, ",")
	for _, v := range vars {
		varVal := strings.Split(v, "=")
		c.register[varVal[0]] = helpers.NaiveStringToInt(varVal[1])
	}
}

func (c Challenge19) RunPartOne(input string) string {
	ps := helpers.SplitParagraphs(input)

	flows := c.parseInput(ps[0])
	lines := helpers.SplitLines(ps[1])

	totalSum := 0

	for _, line := range lines {
		c.setRegister(line)

		state := "in"

		for {
			flow := flows[state]
			matchedRule := false
			for _, r := range flow.rules {
				if c.isTrue(r) {
					state = r.destination
					matchedRule = true
					break
				}
			}
			if !matchedRule {
				state = flow.fallback
			}

			if state == "A" || state == "R" {
				break
			}
		}

		if state == "A" {
			totalSum += c.register["x"] + c.register["m"] + c.register["a"] + c.register["s"]
		}
	}

	return strconv.Itoa(totalSum)
}

func (c Challenge19) RunPartTwo(input string) string {
	ps := helpers.SplitParagraphs(input)
	flows := c.parseInput(ps[0])

	ranges := []rangeContainer{
		{
			state: "in",
			ranges: map[string]aoc_range.Range{
				"x": aoc_range.NewRange(1, 4000, true),
				"m": aoc_range.NewRange(1, 4000, true),
				"a": aoc_range.NewRange(1, 4000, true),
				"s": aoc_range.NewRange(1, 4000, true),
			},
		},
	}

	totalAccepted := 0

	for len(ranges) > 0 {
		rang := ranges[0]
		ranges = ranges[1:]

		if rang.state == "R" {
			continue
		}

		if rang.state == "A" {
			xRange := rang.ranges["x"]
			mRange := rang.ranges["m"]
			aRange := rang.ranges["a"]
			sRange := rang.ranges["s"]
			totalAccepted += xRange.Length() * mRange.Length() * aRange.Length() * sRange.Length()
			continue
		}

		flow := flows[rang.state]
		currentRanges := rang.ranges

		for _, r := range flow.rules {
			varRange := currentRanges[r.variable]

			switch r.operatorStr {
			case ">":
				if varRange.End <= r.value {
					continue
				}

				newMap := maps.Clone(currentRanges)
				newMap[r.variable] = aoc_range.NewRange(r.value+1, varRange.End, true)

				ranges = append(ranges, rangeContainer{
					state:  r.destination,
					ranges: newMap,
				})

				varRange.End = r.value
				currentRanges[r.variable] = varRange

			case "<":
				if varRange.Start >= r.value {
					continue
				}

				newMap := maps.Clone(currentRanges)
				newMap[r.variable] = aoc_range.NewRange(varRange.Start, r.value-1, true)

				ranges = append(ranges, rangeContainer{
					state:  r.destination,
					ranges: newMap,
				})

				varRange.Start = r.value
				currentRanges[r.variable] = varRange

			default:
				panic("this should not happen")
			}

		}

		ranges = append(ranges, rangeContainer{
			state:  flow.fallback,
			ranges: currentRanges,
		})
	}

	return strconv.Itoa(totalAccepted)
}

func (c Challenge19) DataFolder() string {
	return "19"
}
