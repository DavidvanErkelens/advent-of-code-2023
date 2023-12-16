package main

import (
	"advent-of-code-2023/internal/challenges"
	"advent-of-code-2023/internal/clock"
	"advent-of-code-2023/internal/filereader"
	"fmt"
)

func main() {
	allChallenges := []challenges.Challenge{
		challenges.NewChallenge01(),
		challenges.NewChallenge02(),
		challenges.NewChallenge03(),
		challenges.NewChallenge04(),
		//challenges.NewChallenge05(),
		challenges.NewChallenge06(),
		challenges.NewChallenge07(),
		challenges.NewChallenge08(),
		challenges.NewChallenge09(),
		challenges.NewChallenge10(),
		challenges.NewChallenge11(),
		challenges.NewChallenge12(),
		challenges.NewChallenge13(),
		challenges.NewChallenge14(),
		challenges.NewChallenge15(),
		challenges.NewChallenge16(),
	}

	totalTimer := clock.NewClock()
	fr := filereader.InputFileReader()

	for _, challenge := range allChallenges {
		challengeTimer := clock.NewClock()
		input := fr.GetFileDataOrPanic(challenge.DataFolder(), "challenge.in")

		outputOne := challenge.RunPartOne(input)
		fmt.Printf("[%s] Output of part 1 is: %s\n", challenge.DataFolder(), outputOne)

		outputTwo := challenge.RunPartTwo(input)
		fmt.Printf("[%s] Output of part 2 is: %s\n", challenge.DataFolder(), outputTwo)

		challengeTime := challengeTimer.StopClock()
		fmt.Printf("[%s] Runtime: %dms (%dμs)\n\n", challenge.DataFolder(), challengeTime.Milliseconds(), challengeTime.Microseconds())
	}

	totalRuntime := totalTimer.StopClock()
	fmt.Printf("Total runtime: %dms (%dμs)\n", totalRuntime.Milliseconds(), totalRuntime.Microseconds())
}
