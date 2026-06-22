package remove_self_assignments

import "testing"

func TestAddAssign(t *testing.T) {
	// `a += b`→`a = b` だと 3 になり露見する。
	if got := AddAssign(5, 3); got != 8 {
		t.Fatalf("AddAssign(5, 3) = %d, want 8", got)
	}
}
