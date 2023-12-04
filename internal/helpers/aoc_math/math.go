package aoc_math

import "math"

func Abs(input int) int {
	if input < 0 {
		return -1 * input
	}
	return input
}

func Sum(list []int) int {
	total := 0
	for _, value := range list {
		total += value
	}
	return total
}

func Product(list []int) int {
	total := 1
	for _, value := range list {
		total *= value
	}
	return total
}

func IntPow(base, power int) int {
	return int(math.Pow(float64(base), float64(power)))
}
