package helpers

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

func Reduce[TIn any, TOut any](slice []TIn, accumulator func(TIn, TOut) TOut, initial TOut) TOut {
	value := initial

	for _, elem := range slice {
		value = accumulator(elem, value)
	}

	return value
}
