package challenges

import (
	"testing"
)

func TestChallengeCHALLENGE_NUMBER_RunPartOne_ExampleInput(t *testing.T) {
	challenge := NewChallengeCHALLENGE_NUMBER()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.1.out", challenge.RunPartOne)
}

func TestChallengeCHALLENGE_NUMBER_RunPartOne_ChallengeInput(t *testing.T) {
	challenge := NewChallengeCHALLENGE_NUMBER()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.1.out", challenge.RunPartOne)
}

func TestChallengeCHALLENGE_NUMBER_RunPartTwo_ExampleInput(t *testing.T) {
	challenge := NewChallengeCHALLENGE_NUMBER()
	RunChallengeTest(t, challenge.DataFolder(), "example.in", "example.2.out", challenge.RunPartTwo)
}

func TestChallengeCHALLENGE_NUMBER_RunPartTwo_ChallengeInput(t *testing.T) {
	challenge := NewChallengeCHALLENGE_NUMBER()
	RunChallengeTest(t, challenge.DataFolder(), "challenge.in", "challenge.2.out", challenge.RunPartTwo)
}
