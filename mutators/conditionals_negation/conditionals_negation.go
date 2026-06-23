// Package conditionals_negation demonstrates the CONDITIONALS_NEGATION mutator (e.g. == -> !=).
package conditionals_negation

// IsZero returns n == 0. gremlins mutates `==` into `!=`.
func IsZero(n int) bool {
	return n == 0
}
