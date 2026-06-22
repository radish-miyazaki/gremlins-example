// Package invert_bitwise は INVERT_BITWISE mutator（& → | など、既定無効）のデモ。
package invert_bitwise

// And は a & b を返す。gremlins は `&` を `|` に変異させる。
func And(a, b int) int {
	return a & b
}
