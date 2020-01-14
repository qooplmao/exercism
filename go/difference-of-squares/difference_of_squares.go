package diffsquares

/**

Calculate the square of the sum of all integers up to a limit
eg.  (1 + 2 + 3)² = 6² = 36

*/
func SquareOfSum(highest int) int {
	sum := 0

	for i := 1; i <= highest; i += 1 {
		sum += i
	}

	return sum *  sum
}

/**

Calculate the sum of squares of all integers up to a limit
eg. 1² + 2² + 3² = 1 + 4 + 9 = 14

*/
func SumOfSquares(highest int) int {
	sum := 0

	for i := 1; i <= highest; i += 1 {
		sum += i * i
	}

	return sum
}

/**

Calculate the difference between the "square of sum" and "sum of squares"
eg. (1 + 2 + 3)² - (1² + 2² + 3²) = 36 - 14 = 22

*/
func Difference(i int) int {
	return SquareOfSum(i) - SumOfSquares(i)
}
