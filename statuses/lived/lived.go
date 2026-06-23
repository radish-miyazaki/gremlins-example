// Package lived demonstrates the LIVED status (a weak test that misses the mutation).
package lived

// Double returns n + n. gremlins mutates `+` into `-`, but because the test only
// checks Double(0), the mutation survives since 0-0==0 (LIVED).
func Double(n int) int {
	return n + n
}
