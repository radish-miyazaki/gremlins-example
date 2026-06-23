// Package conditionals_boundary demonstrates the CONDITIONALS_BOUNDARY mutator (e.g. < -> <=).
package conditionals_boundary

// BelowLimit reports whether n is below 10. gremlins mutates `<` into `<=`.
func BelowLimit(n int) bool {
	return n < 10
}
