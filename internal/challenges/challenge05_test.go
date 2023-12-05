package challenges

import (
	"testing"
)

func TestChallenge05_RunPartOne_ExampleInput(t *testing.T) {
	challenge := NewChallenge05()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.1.out", challenge.RunPartOne)
}

func TestChallenge05_RunPartOne_ChallengeInput(t *testing.T) {
	challenge := NewChallenge05()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.1.out", challenge.RunPartOne)
}

func TestChallenge05_RunPartTwo_ExampleInput(t *testing.T) {
	challenge := NewChallenge05()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.2.out", challenge.RunPartTwo)
}

func TestChallenge05_RunPartTwo_ChallengeInput(t *testing.T) {
	challenge := NewChallenge05()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.2.out", challenge.RunPartTwo)
}
