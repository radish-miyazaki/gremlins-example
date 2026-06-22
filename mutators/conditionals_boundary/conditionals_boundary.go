// Package conditionals_boundary は CONDITIONALS_BOUNDARY mutator（< → <= など）のデモ。
package conditionals_boundary

// BelowLimit は n が 10 未満かを返す。gremlins は `<` を `<=` に変異させる。
func BelowLimit(n int) bool {
	return n < 10
}
