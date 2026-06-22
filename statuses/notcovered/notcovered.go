// Package notcovered は NOT COVERED ステータスのデモ。
// Covered はテストされ（RUNNABLE/KILLED）、Uncovered はテストされない（NOT COVERED）。
package notcovered

// Covered はテストでカバーされる。
func Covered(a, b int) int {
	return a + b
}

// Uncovered はテストから呼ばれないため、その `+` の変異は NOT COVERED になる。
func Uncovered(a, b int) int {
	return a + b
}
