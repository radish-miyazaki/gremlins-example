// Package remove_self_assignments demonstrates the REMOVE_SELF_ASSIGNMENTS mutator (a += b -> a = b, disabled by default).
package remove_self_assignments

// AddAssign returns the result of a += b. gremlins mutates the compound assignment `+=` into a plain assignment `=`.
func AddAssign(a, b int) int {
	a += b
	return a
}
