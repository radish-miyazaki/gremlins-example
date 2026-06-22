// Package timedout は TIMED OUT ステータスのデモ。
// ループの `i++` を `i--` に変異させると無限ループになり、テストがタイムアウトする。
package timedout

// SumTo は 0 から n-1 までの総和を返す。gremlins が `i++` を `i--` に
// 変異させると i は減り続け i < n が永遠に真となり無限ループ（TIMED OUT）。
func SumTo(n int) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += i
	}
	return sum
}
