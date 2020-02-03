package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency returns the frequency of each rune given an array of strings
// and return as a FreqMap.
func ConcurrentFrequency(slice []string) FreqMap {
	// create a buffered channel with capacity equal to the slice length
	c := make(chan FreqMap, 10)
	for _, r := range slice {
		r := r
		go func() {
			c <- Frequency(r)
		}() // spin up multiple go-routines to calculate the frequency
	}

	m := FreqMap{} // returned frequency map
	for range slice {
		for word, frequency := range <-c {
			m[word] += frequency
		}
	}
	return m
}
