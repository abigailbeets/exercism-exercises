package triangle

import (
	"math"
)

const testVersion = 3

type Kind string

const NaT = "Not a triangle"
const Equ = "Equilateral"
const Iso = "Isosceles"
const Sca = "Scalene"

func KindFromSides(a, b, c float64) Kind {
	inputArray := []float64{a, b, c}

	switch {
	case any(inputArray, lessThanOne) || (a+b < c) || (a+c < b) || (c+b < a):
		return NaT
	case a == b && a == c && !any(inputArray, lessThanOne):
		return Equ
	case a == b || a == c || b == c && !any(inputArray, lessThanOne):
		return Iso
	default:
		return Sca
	}
}

func lessThanOne(number float64) bool {
	if number <= 0 || math.IsNaN(number) || math.IsInf(number, 1) || math.IsInf(number, -1) {
		return true
	}
	return false
}

func any(input []float64, myFunc func(float64) bool) bool {
	for _, item := range input {
		if myFunc(item) {
			return true
		}
	}
	return false
}
