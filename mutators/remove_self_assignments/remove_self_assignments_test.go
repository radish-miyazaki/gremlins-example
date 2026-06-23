package remove_self_assignments

import "testing"

func TestAddAssign(t *testing.T) {
	// With `a += b`->`a = b` it becomes 3, exposing the mutation.
	if got := AddAssign(5, 3); got != 8 {
		t.Fatalf("AddAssign(5, 3) = %d, want 8", got)
	}
}
