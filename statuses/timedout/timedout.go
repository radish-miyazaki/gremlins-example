// Package timedout demonstrates the TIMED OUT status.
// Mutating the loop's `i++` into `i--` causes an infinite loop, so the test times out.
package timedout

// SumTo returns the sum from 0 to n-1. When gremlins mutates `i++` into `i--`,
// i keeps decreasing so i < n stays true forever, causing an infinite loop (TIMED OUT).
func SumTo(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
