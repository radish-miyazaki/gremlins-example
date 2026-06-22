package invert_assignments

import "testing"

func TestAddAssign(t *testing.T) {
	if got := AddAssign(5, 3); got != 8 {
		t.Fatalf("AddAssign(5, 3) = %d, want 8", got)
	}
}
