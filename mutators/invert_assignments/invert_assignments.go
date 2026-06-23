// Package invert_assignments demonstrates the INVERT_ASSIGNMENTS mutator (e.g. += -> -=, disabled by default).
package invert_assignments

// AddAssign returns the result of a += b. gremlins mutates `+=` into `-=`.
func AddAssign(a, b int) int {
	a += b
	return a
}
