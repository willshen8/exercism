package darts

import "math"

// Score determines the score given the x and y coordinate of the point of contact
func Score(x, y float64) int {
	radius := math.Sqrt(x*x + y*y)
	if radius > 10 {
		return 0
	} else if radius > 5 && radius <= 10 {
		return 1
	} else if radius <= 5 && radius > 1 {
		return 5
	}
	return 10
}
