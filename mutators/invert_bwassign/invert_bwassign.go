// Package invert_bwassign は INVERT_BWASSIGN mutator（&= → |= など、既定無効）のデモ。
package invert_bwassign

// AndAssign は a &= b の結果を返す。gremlins は `&=` を `|=` に変異させる。
func AndAssign(a, b int) int {
	a &= b
	return a
}
