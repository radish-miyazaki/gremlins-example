package lived

import "testing"

// A deliberately weak test that only checks n=0. It cannot kill the `+`->`-` mutation, so the result is LIVED.
func TestDoubleWeak(t *testing.T) {
	if got := Double(0); got != 0 {
		t.Fatalf("Double(0) = %d, want 0", got)
	}
}
