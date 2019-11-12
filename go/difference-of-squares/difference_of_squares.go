// Package diffsquares gives the difference of SquareOfSum and SquareOfSum
package diffsquares

// SquareOfSum returns the square of the sum of the first N natural numbers.
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// SquareOfSum returns the sum of the first N natural numbers squared.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference of a and b
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
