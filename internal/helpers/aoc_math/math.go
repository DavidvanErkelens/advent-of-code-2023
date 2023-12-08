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

// GCD and LCM borrowed from https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(values ...int) int {
	a := values[0]
	b := values[1]
	rest := values[2:]
	result := a * b / GCD(a, b)

	for i := 0; i < len(rest); i++ {
		result = LCM(result, rest[i])
	}

	return result
}
