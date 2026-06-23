package invert_loopctrl

import "testing"

func TestFirstHit(t *testing.T) {
	// Returns the index of the first true. With `break`->`continue` it returns the last true (=2), exposing the mutation.
	if got := FirstHit([]bool{false, true, true}); got != 1 {
		t.Fatalf("FirstHit({false,true,true}) = %d, want 1", got)
	}
}
