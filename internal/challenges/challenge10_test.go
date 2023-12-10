package challenges

import (
	"testing"
)

func TestChallenge10_RunPartOne_ExampleInput(t *testing.T) {
	challenge := NewChallenge10()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.1.out", challenge.RunPartOne)
}

func TestChallenge10_RunPartOne_ChallengeInput(t *testing.T) {
	challenge := NewChallenge10()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.1.out", challenge.RunPartOne)
}

func TestChallenge10_RunPartTwo_ExampleInput(t *testing.T) {
	challenge := NewChallenge10()
	RunChallengeTest(t, challenge.DataFolder(), "example.2.in", "example.2.out", challenge.RunPartTwo)
}

func TestChallenge10_RunPartTwo_ChallengeInput(t *testing.T) {
	challenge := NewChallenge10()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.2.out", challenge.RunPartTwo)
}
