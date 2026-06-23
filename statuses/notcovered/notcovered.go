// Package notcovered demonstrates the NOT COVERED status.
// Covered is exercised by a test (RUNNABLE/KILLED); Uncovered is not (NOT COVERED).
package notcovered

// Covered is exercised by a test.
func Covered(a, b int) int {
	return a + b
}

// Uncovered is never called from a test, so the mutation of its `+` is NOT COVERED.
func Uncovered(a, b int) int {
	return a + b
}
