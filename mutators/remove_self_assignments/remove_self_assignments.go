// Package remove_self_assignments は REMOVE_SELF_ASSIGNMENTS mutator（a += b → a = b、既定無効）のデモ。
package remove_self_assignments

// AddAssign は a += b の結果を返す。gremlins は複合代入 `+=` を単純代入 `=` に変異させる。
func AddAssign(a, b int) int {
	a += b
	return a
}
