package invert_bitwise

import "testing"

func TestAnd(t *testing.T) {
	// 6 (110) & 3 (011) = 2 (010). With `&`->`|` it becomes 7 (111), exposing the mutation.
	if got := And(6, 3); got != 2 {
		t.Fatalf("And(6, 3) = %d, want 2", got)
	}
}
