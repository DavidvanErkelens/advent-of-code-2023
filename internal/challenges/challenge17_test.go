package challenges

import (
	"testing"
)

func TestChallenge17_RunPartOne_ExampleInput(t *testing.T) {
	challenge := NewChallenge17()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.1.out", challenge.RunPartOne)
}

func TestChallenge17_RunPartOne_ChallengeInput(t *testing.T) {
	challenge := NewChallenge17()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.1.out", challenge.RunPartOne)
}

func TestChallenge17_RunPartTwo_ExampleInput(t *testing.T) {
	challenge := NewChallenge17()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.2.out", challenge.RunPartTwo)
}

func TestChallenge17_RunPartTwo_ChallengeInput(t *testing.T) {
	challenge := NewChallenge17()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.2.out", challenge.RunPartTwo)
}
