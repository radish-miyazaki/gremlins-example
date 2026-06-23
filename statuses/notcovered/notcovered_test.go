package notcovered

import "testing"

func TestCovered(t *testing.T) {
	if got := Covered(2, 3); got != 5 {
		t.Fatalf("Covered(2, 3) = %d, want 5", got)
	}
}
