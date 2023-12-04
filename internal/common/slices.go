package common

func NewSlice(start, end, step int) []int {
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

func Filter[T any](slice []T, filter func(T) bool) []T {
	out := make([]T, 0)

	for _, elem := range slice {
		if filter(elem) {
			out = append(out, elem)
		}
	}

	return out
}

func Map[TIn any, TOut any](slice []TIn, mapper func(TIn) TOut) []TOut {
	out := make([]TOut, 0)

	for _, elem := range slice {
		out = append(out, mapper(elem))
	}

	return out
}
