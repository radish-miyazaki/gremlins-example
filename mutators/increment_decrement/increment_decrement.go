// Package increment_decrement demonstrates the INCREMENT_DECREMENT mutator (++ -> --).
package increment_decrement

// Next increments n and returns it. gremlins mutates `++` into `--`.
func Next(n int) int {
	n++
	return n
}
