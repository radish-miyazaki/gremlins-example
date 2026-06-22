// Package arithmetic_base は ARITHMETIC_BASE mutator（+ → - など）のデモ。
package arithmetic_base

// Add は a + b を返す。gremlins は `+` を `-` に変異させる。
func Add(a, b int) int {
	return a + b
}
