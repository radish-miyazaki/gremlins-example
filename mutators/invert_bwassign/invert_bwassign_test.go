package invert_bwassign

import "testing"

func TestAndAssign(t *testing.T) {
	// 6 &= 3 -> 2. With `&=`->`|=` it becomes 7, exposing the mutation.
	if got := AndAssign(6, 3); got != 2 {
		t.Fatalf("AndAssign(6, 3) = %d, want 2", got)
	}
}
