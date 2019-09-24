// Package triangle determines if a triangle is equilateral, isosceles, or scalene.
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a == b && b == c {
		return Equ
	}
	if a != b && b != c && a != c {
		return Sca
	}
	if a == b || a == c || b == c {
		return Iso
	}

	return NaT
}
