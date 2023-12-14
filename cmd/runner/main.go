package main

import (
	"advent-of-code-2023/internal/challenges"
	"advent-of-code-2023/internal/clock"
	"advent-of-code-2023/internal/filereader"
	"fmt"
)

func main() {
	challenge := challenges.NewChallenge14()
	fr := filereader.InputFileReader()
	input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")

	partOneClock := clock.NewClock()
	outputOne := challenge.RunPartOne(input)
	partOneRuntime := partOneClock.StopClock()
	fmt.Printf("Output of part 1 is: %s\n", outputOne)
	fmt.Printf("Runtime: %dms (%dμs)\n\n", partOneRuntime.Milliseconds(), partOneRuntime.Microseconds())

	partTwoClock := clock.NewClock()
	outputTwo := challenge.RunPartTwo(input)
	partTwoRuntime := partTwoClock.StopClock()
	fmt.Printf("Output of part 2 is: %s\n", outputTwo)
	fmt.Printf("Runtime: %dms (%dμs)\n\n", partTwoRuntime.Milliseconds(), partTwoRuntime.Microseconds())

	totalRuntime := partOneClock.StopClock()
	fmt.Printf("Total runtime: %dms (%dμs)\n", totalRuntime.Milliseconds(), totalRuntime.Microseconds())
}
