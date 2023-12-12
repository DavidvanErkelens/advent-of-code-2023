package helpers

import (
	"strconv"
	"strings"
)

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

func ReduceWithIndex[TIn any, TOut any](slice []TIn, accumulator func(int, TIn, TOut) TOut, initial TOut) TOut {
	value := initial

	for idx, elem := range slice {
		value = accumulator(idx, elem, value)
	}

	return value
}

func GetLastElement[T any](s []T) T {
	return s[len(s)-1]
}

func ContainsElement[T comparable](s []T, v T) bool {
	for _, elem := range s {
		if elem == v {
			return true
		}
	}
	return false
}

func AllSatisfies[T comparable](s []T, check func(T) bool) bool {
	for _, v := range s {
		if !check(v) {
			return false
		}
	}

	return true
}

func Insert[T any](a []T, index int, value T) []T {
	if len(a) == index {
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...)
	a[index] = value
	return a
}

func ToString(s []int) string {
	return strings.Join(Map(s, func(in int) string {
		return strconv.Itoa(in)
	}), ",")
}
