// Package invert_negatives demonstrates the INVERT_NEGATIVES mutator (unary -x -> +x).
package invert_negatives

// Negate returns -n. gremlins inverts the unary `-` (effectively removing it).
func Negate(n int) int {
	return -n
}
