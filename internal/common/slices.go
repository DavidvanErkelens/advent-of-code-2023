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
