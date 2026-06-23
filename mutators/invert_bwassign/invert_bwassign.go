// Package invert_bwassign demonstrates the INVERT_BWASSIGN mutator (e.g. &= -> |=, disabled by default).
package invert_bwassign

// AndAssign returns the result of a &= b. gremlins mutates `&=` into `|=`.
func AndAssign(a, b int) int {
	a &= b
	return a
}
