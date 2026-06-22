package invert_bitwise

import "testing"

func TestAnd(t *testing.T) {
	// 6 (110) & 3 (011) = 2 (010)。`&`→`|` だと 7 (111) になり露見する。
	if got := And(6, 3); got != 2 {
		t.Fatalf("And(6, 3) = %d, want 2", got)
	}
}
