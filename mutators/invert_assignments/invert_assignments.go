// Package invert_assignments は INVERT_ASSIGNMENTS mutator（+= → -= など、既定無効）のデモ。
package invert_assignments

// AddAssign は a += b の結果を返す。gremlins は `+=` を `-=` に変異させる。
func AddAssign(a, b int) int {
	a += b
	return a
}
