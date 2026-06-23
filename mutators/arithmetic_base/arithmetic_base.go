// Package arithmetic_base demonstrates the ARITHMETIC_BASE mutator (e.g. + -> -).
package arithmetic_base

// Add returns a + b. gremlins mutates `+` into `-`.
func Add(a, b int) int {
	return a + b
}
