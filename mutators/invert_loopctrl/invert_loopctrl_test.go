package invert_loopctrl

import "testing"

func TestFirstHit(t *testing.T) {
	// 最初に true になる添字を返す。`break`→`continue` だと最後の true (=2) を返し露見する。
	if got := FirstHit([]bool{false, true, true}); got != 1 {
		t.Fatalf("FirstHit({false,true,true}) = %d, want 1", got)
	}
}
