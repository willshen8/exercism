// Package secret contains the code to decode secret handshake
package secret

// Handshake converts an integer into a series of events
func Handshake(n uint) []string {
	var secretMessage []string
	if n&uint(1) == 1 {
		secretMessage = append(secretMessage, "wink")
	}
	if n&uint(2) == 2 {
		secretMessage = append(secretMessage, "double blink")
	}
	if n&uint(4) == 4 {
		secretMessage = append(secretMessage, "close your eyes")
	}
	if n&uint(8) == 8 {
		secretMessage = append(secretMessage, "jump")
	}
	if n&uint(16) == 16 {
		secretMessage = Reverse(secretMessage)
	}
	return secretMessage
}

// Reverse returns the string reversed
func Reverse(input []string) []string {
	var reversed []string
	for i := len(input) - 1; i > -1; i-- {
		reversed = append(reversed, input[i])
	}
	return reversed
}
