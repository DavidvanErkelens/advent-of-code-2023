package aoc_loop

type Loop struct {
	First int
	Last  int
}

func NewLoop(first, last int) Loop {
	return Loop{First: first, Last: last}
}

func (l Loop) FindLastEqual(end int) int {
	loopSize := l.Last - l.First
	stepsFromLast := end - l.Last
	return end - (stepsFromLast % loopSize)
}

func (l Loop) FindEquivalent(value int) int {
	loopSize := l.Last - l.First
	stepsFromLast := value - l.Last
	return l.First + (stepsFromLast % loopSize)
}
