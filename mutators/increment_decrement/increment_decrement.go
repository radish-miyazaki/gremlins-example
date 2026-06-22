// Package increment_decrement は INCREMENT_DECREMENT mutator（++ → --）のデモ。
package increment_decrement

// Next は n をインクリメントして返す。gremlins は `++` を `--` に変異させる。
func Next(n int) int {
	n++
	return n
}
