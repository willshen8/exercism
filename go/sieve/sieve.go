package sieve

import (
	"sort"
)

func Sieve(limit int) []int {
	var output []int
	numbers := make(map[int]bool)
	if limit < 2 {
		return output
	}

	for i := 2; i <= limit; i++ {
		numbers[i] = true
	}

	for i := 2; i <= limit; i = NextPrime(i) {
		for j := 2; i*j <= limit; j++ {
			delete(numbers, i*j)
		}
	}

	for prime := range numbers {
		output = append(output, prime)
	}
	sort.Ints(output)
	return output
}

// NextPrime calculates the next prime given the last prime calculated
func NextPrime(curPrime int) int {
	var nextPrime int
	for i := curPrime + 1; nextPrime == 0; i++ {
		if IsPrime(i) {
			nextPrime = i
		}
	}
	return nextPrime
}

// IsPrime test to see if a number is a prime number
func IsPrime(n int) bool {
	for i := 2; i < n-1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
