// Package invert_logical demonstrates the INVERT_LOGICAL mutator (&& -> ||, disabled by default).
package invert_logical

// Both returns a && b. gremlins mutates `&&` into `||`.
func Both(a, b bool) bool {
	return a && b
}
