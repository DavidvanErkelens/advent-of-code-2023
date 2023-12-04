package aoc_range

func NewRange(start, end int) []int {
	return NewRangeWithStep(start, end, 1)
}

func NewRangeWithStep(start, end, step int) []int {
	if step <= 0 || end < start {
		return []int{}
	}
	s := make([]int, 0, 1+(end-start)/step)
	for start <= end {
		s = append(s, start)
		start += step
	}
	return s
}
