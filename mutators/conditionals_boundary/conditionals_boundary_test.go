package conditionals_boundary

import "testing"

func TestBelowLimit(t *testing.T) {
	if !BelowLimit(9) {
		t.Fatalf("BelowLimit(9) = false, want true")
	}
	// Boundary value: 10 is not below the limit. The `<`->`<=` mutation is exposed by this case.
	if BelowLimit(10) {
		t.Fatalf("BelowLimit(10) = true, want false")
	}
}
