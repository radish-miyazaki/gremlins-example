// Package invert_logical は INVERT_LOGICAL mutator（&& → ||、既定無効）のデモ。
package invert_logical

// Both は a && b を返す。gremlins は `&&` を `||` に変異させる。
func Both(a, b bool) bool {
	return a && b
}
