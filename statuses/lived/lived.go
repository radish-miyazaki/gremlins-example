// Package lived は LIVED ステータスのデモ（弱いテストが変異を見逃すケース）。
package lived

// Double は n + n を返す。gremlins は `+` を `-` に変異させるが、
// テストが Double(0) しか検証しないため 0-0==0 で生き残る（LIVED）。
func Double(n int) int {
	return n + n
}
