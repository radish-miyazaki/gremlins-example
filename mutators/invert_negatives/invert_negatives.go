// Package invert_negatives は INVERT_NEGATIVES mutator（単項 -x → +x）のデモ。
package invert_negatives

// Negate は -n を返す。gremlins は単項 `-` を反転（実質除去）する。
func Negate(n int) int {
	return -n
}
