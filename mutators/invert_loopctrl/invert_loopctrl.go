// Package invert_loopctrl demonstrates the INVERT_LOOPCTRL mutator (break -> continue, disabled by default).
package invert_loopctrl

// FirstHit returns the index of the first true value in flags (-1 if there is none).
// gremlins mutates `break` into `continue`.
func FirstHit(flags []bool) int {
	idx := -1
	for i, f := range flags {
		if f {
			idx = i
			break
		}
	}
	return idx
}
