package triangle

import "math"

const testVersion = 3

type Kind string

const (
	Equ Kind = "equilateral"
	Iso Kind = "isosceles"
	Sca Kind = "scalene"
	NaT Kind = "not a triangle"
)

// Determines if the sides create an existent triangle
func isValidTriangle(a, b, c float64) bool {
	sides := []float64{a, b, c}
	// Make sure all sides are numbers and not infinity
	for _, x := range sides {
		if math.IsNaN(x) || math.IsInf(x, 0) {
			return false
		}
	}
	// Test triangle ineequality theorm
	if ((a + b >= c) && (b + c >= a) && (a + c >= b)) {
		return true
	}
	return false
}

// Returns the kind of triangle
func KindFromSides(a, b, c float64) Kind {
	// If sides are 0 or less, not a triangle
  if (a <= 0 || b <= 0 || c <= 0) {
		return NaT
	}
	isValid := isValidTriangle(a, b, c)
	if !isValid {
		return NaT
	}
	if ( a == b && b == c) {
		return Equ
	}
	if (a == b || b == c || a == c) {
		return Iso
	}
	return Sca
}
