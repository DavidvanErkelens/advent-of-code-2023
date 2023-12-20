package challenges

import (
	"advent-of-code-2023/internal/helpers"
	"advent-of-code-2023/internal/helpers/aoc_math"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge20() Challenge20 {
	return Challenge20{}
}

type Challenge20 struct {
}

const (
	flipFlop    = "%"
	conjunction = "&"
)

const (
	highPulse = 1
	lowPulse  = 0
)

type module struct {
	operator     string
	destinations []string
}

type pulse struct {
	from  string
	to    string
	pulse int
}

func (c Challenge20) parseInput(input string) map[string]module {
	lines := helpers.SplitLines(input)

	modules := make(map[string]module)

	for _, l := range lines {
		data := strings.Split(l, " -> ")
		name := strings.TrimLeft(data[0], "%&")
		destinations := helpers.Map(strings.Split(data[1], ","), func(in string) string { return strings.TrimSpace(in) })

		modules[name] = module{
			operator:     string(l[0]),
			destinations: destinations,
		}
	}

	return modules
}

func (c Challenge20) RunPartOne(input string) string {
	modules := c.parseInput(input)
	conjectionMemory := make(map[string]map[string]int)
	flipFlopState := make(map[string]int)

	for name, m := range modules {
		if m.operator == conjunction {
			conjectionMemory[name] = make(map[string]int)

			for name2, m2 := range modules {
				if helpers.ContainsElement(m2.destinations, name) {
					conjectionMemory[name][name2] = lowPulse
				}
			}
		}

		if m.operator == flipFlop {
			flipFlopState[name] = lowPulse
		}
	}

	highPulses := 0
	lowPulses := 0

	for i := 0; i < 1000; i++ {

		pulses := []pulse{
			{
				from:  "button",
				to:    "broadcaster",
				pulse: 0,
			},
		}

		for len(pulses) > 0 {

			p := pulses[0]
			pulses = pulses[1:]
			//fmt.Println(p.from, "--", p.pulse, "->", p.to)
			if p.pulse == lowPulse {
				lowPulses += 1
			} else {
				highPulses += 1
			}

			targetModule, ok := modules[p.to]
			if !ok {
				//fmt.Println("Target not found:", p.to)
				continue
			}
			if targetModule.operator == flipFlop {
				if p.pulse == highPulse {
					continue
				}

				flipFlopState[p.to] = 1 - flipFlopState[p.to]
				for _, dest := range targetModule.destinations {
					pulses = append(pulses, pulse{
						from:  p.to,
						to:    dest,
						pulse: flipFlopState[p.to],
					})
				}
			}

			if targetModule.operator == conjunction {
				conjectionMemory[p.to][p.from] = p.pulse
				pulseToSend := lowPulse
				for _, connected := range conjectionMemory[p.to] {
					if connected == lowPulse {
						pulseToSend = highPulse
						break
					}
				}
				for _, dest := range targetModule.destinations {
					pulses = append(pulses, pulse{
						from:  p.to,
						to:    dest,
						pulse: pulseToSend,
					})
				}
			}

			// start case
			if targetModule.operator == "b" {
				for _, dest := range targetModule.destinations {
					pulses = append(pulses, pulse{
						from:  p.to,
						to:    dest,
						pulse: lowPulse,
					})
				}
			}
		}
	}
	return strconv.Itoa(lowPulses * highPulses)
}

func (c Challenge20) RunPartTwo(input string) string {
	modules := c.parseInput(input)

	// no example case here
	if len(modules) < 10 {
		return "0"
	}

	conjectionMemory := make(map[string]map[string]int)
	flipFlopState := make(map[string]int)

	for name, m := range modules {
		if m.operator == conjunction {
			conjectionMemory[name] = make(map[string]int)

			for name2, m2 := range modules {
				if helpers.ContainsElement(m2.destinations, name) {
					conjectionMemory[name][name2] = lowPulse
				}
			}
		}

		if m.operator == flipFlop {
			flipFlopState[name] = lowPulse
		}
	}

	i := 0

	conjuctionBefore := "mg"
	required := make([]string, 0)

	for key := range conjectionMemory[conjuctionBefore] {
		required = append(required, key)
	}

	values := make([]int, 0)

	for {

		pulses := []pulse{
			{
				from:  "button",
				to:    "broadcaster",
				pulse: 0,
			},
		}

		for len(pulses) > 0 {

			p := pulses[0]
			pulses = pulses[1:]
			//fmt.Println(p.from, "--", p.pulse, "->", p.to)

			if p.to == "mg" && p.pulse == highPulse && helpers.ContainsElement(required, p.from) {
				values = append(values, i+1)
				required = helpers.RemoveIndex(required, slices.Index(required, p.from))

				if len(required) == 0 {
					break
				}
			}

			targetModule, ok := modules[p.to]
			if !ok {
				//fmt.Println("Target not found:", p.to)
				continue
			}
			if targetModule.operator == flipFlop {
				if p.pulse == highPulse {
					continue
				}

				flipFlopState[p.to] = 1 - flipFlopState[p.to]
				for _, dest := range targetModule.destinations {
					pulses = append(pulses, pulse{
						from:  p.to,
						to:    dest,
						pulse: flipFlopState[p.to],
					})
				}
			}

			if targetModule.operator == conjunction {
				conjectionMemory[p.to][p.from] = p.pulse
				pulseToSend := lowPulse
				for _, connected := range conjectionMemory[p.to] {
					if connected == lowPulse {
						pulseToSend = highPulse
						break
					}
				}
				for _, dest := range targetModule.destinations {
					pulses = append(pulses, pulse{
						from:  p.to,
						to:    dest,
						pulse: pulseToSend,
					})
				}
			}

			// start case
			if targetModule.operator == "b" {
				for _, dest := range targetModule.destinations {
					pulses = append(pulses, pulse{
						from:  p.to,
						to:    dest,
						pulse: lowPulse,
					})
				}
			}
		}

		if len(required) == 0 {
			break
		}

		i += 1
	}
	return strconv.Itoa(aoc_math.LCM(values...))
}

func (c Challenge20) DataFolder() string {
	return "20"
}
