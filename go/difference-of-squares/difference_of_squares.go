package diffsquares

const testVersion = 1

func SquareOfSums(startingNum int) int {
	result := 0
	for i := startingNum; i > 0; i-- {
		result += i
	}

	result *= result
	return result
}

func SumOfSquares(startingNum int) int {
	result := 0
	for i := startingNum; i > 0; i-- {
		result += i * i

	}

	return result
}

func Difference(startingNum int) int {
	return SquareOfSums(startingNum) - SumOfSquares(startingNum)
}
