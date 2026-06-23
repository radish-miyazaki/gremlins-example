// Package invert_bitwise demonstrates the INVERT_BITWISE mutator (e.g. & -> |, disabled by default).
package invert_bitwise

// And returns a & b. gremlins mutates `&` into `|`.
func And(a, b int) int {
	return a & b
}
