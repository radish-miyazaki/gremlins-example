package invert_negatives

import "testing"

func TestNegate(t *testing.T) {
	if got := Negate(5); got != -5 {
		t.Fatalf("Negate(5) = %d, want -5", got)
	}
}
