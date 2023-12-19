package aoc_range

func NewRangeSlice(start, end int) []int {
	return NewRangeSliceWithStep(start, end, 1)
}

func NewRangeSliceWithStep(start, end, step int) []int {
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

func NewRange(Start, End int, inclusiveEnd bool) Range {
	return Range{
		Start:        Start,
		End:          End,
		inclusiveEnd: inclusiveEnd,
	}
}

type Range struct {
	Start        int
	End          int
	inclusiveEnd bool
}

func (r Range) Length() int {
	if r.inclusiveEnd {
		return r.End - r.Start + 1
	}
	return r.End - r.Start
}
