// Package diffsquares gives the difference of SquareOfSum and SquareOfSum
package diffsquares

// SquareOfSum returns the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i
	}
	return sum * sum
}

// SquareOfSum returns the sum of the first N natural numbers squared.
func SumOfSquares(n int) int {
	var sum int
	for i := n; i > 0; i-- {
		sum += i * i
	}
	return sum
}

// Difference returns the difference of a and b
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
