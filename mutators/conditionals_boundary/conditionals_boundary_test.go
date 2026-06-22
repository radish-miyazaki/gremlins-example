package conditionals_boundary

import "testing"

func TestBelowLimit(t *testing.T) {
	if !BelowLimit(9) {
		t.Fatalf("BelowLimit(9) = false, want true")
	}
	// 境界値: 10 は限界未満ではない。`<`→`<=` の変異はこのケースで露見する。
	if BelowLimit(10) {
		t.Fatalf("BelowLimit(10) = true, want false")
	}
}
