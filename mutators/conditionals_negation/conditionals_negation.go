// Package conditionals_negation は CONDITIONALS_NEGATION mutator（== → != など）のデモ。
package conditionals_negation

// IsZero は n == 0 を返す。gremlins は `==` を `!=` に変異させる。
func IsZero(n int) bool {
	return n == 0
}
