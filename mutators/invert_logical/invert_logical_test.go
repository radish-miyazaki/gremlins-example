package invert_logical

import "testing"

func TestBoth(t *testing.T) {
	if Both(true, true) != true {
		t.Fatalf("Both(true, true) = false, want true")
	}
	// `&&`→`||` だと Both(true, false) が true になり露見する。
	if Both(true, false) != false {
		t.Fatalf("Both(true, false) = true, want false")
	}
}
