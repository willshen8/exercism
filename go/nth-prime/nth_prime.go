package prime

// Nth returns the Nth prime number
func Nth(n int) (int, bool) {
	if n == 0 {
		return 0, false
	}

	primes := []int{2}

	for len(primes) < n {
		primes = append(primes, NextPrime(primes[len(primes)-1]))
	}
	return primes[n-1], true
}

// NextPrime calculates the next prime given the last prime calculated
func NextPrime(curPrime int) int {
	var nextPrime int
	for i := curPrime + 1; nextPrime == 0; i++ {
		if IsPrime(i) == true {
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
