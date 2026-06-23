package increment_decrement

import "testing"

func TestNext(t *testing.T) {
	if got := Next(5); got != 6 {
		t.Fatalf("Next(5) = %d, want 6", got)
	}
}
