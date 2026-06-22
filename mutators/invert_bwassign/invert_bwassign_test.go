package invert_bwassign

import "testing"

func TestAndAssign(t *testing.T) {
	// 6 &= 3 -> 2。`&=`→`|=` だと 7 になり露見する。
	if got := AndAssign(6, 3); got != 2 {
		t.Fatalf("AndAssign(6, 3) = %d, want 2", got)
	}
}
