package lived

import "testing"

// あえて n=0 だけを検証する弱いテスト。`+`→`-` の変異を殺せず LIVED になる。
func TestDoubleWeak(t *testing.T) {
	if got := Double(0); got != 0 {
		t.Fatalf("Double(0) = %d, want 0", got)
	}
}
