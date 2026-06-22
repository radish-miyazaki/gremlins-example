package conditionals_negation

import "testing"

func TestIsZero(t *testing.T) {
	if !IsZero(0) {
		t.Fatalf("IsZero(0) = false, want true")
	}
	if IsZero(1) {
		t.Fatalf("IsZero(1) = true, want false")
	}
}
