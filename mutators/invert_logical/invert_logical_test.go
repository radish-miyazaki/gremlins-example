package invert_logical

import "testing"

func TestBoth(t *testing.T) {
	if Both(true, true) != true {
		t.Fatalf("Both(true, true) = false, want true")
	}
	// With `&&`->`||`, Both(true, false) returns true, exposing the mutation.
	if Both(true, false) != false {
		t.Fatalf("Both(true, false) = true, want false")
	}
}
