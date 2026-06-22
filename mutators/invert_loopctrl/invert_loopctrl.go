// Package invert_loopctrl は INVERT_LOOPCTRL mutator（break → continue、既定無効）のデモ。
package invert_loopctrl

// FirstHit は flags の中で最初に true になる添字を返す（無ければ -1）。
// gremlins は `break` を `continue` に変異させる。
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
