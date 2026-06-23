package timedout

import "testing"

func TestSumTo(t *testing.T) {
	// 0+1+2+3+4 = 10
	if got := SumTo(5); got != 10 {
		t.Fatalf("SumTo(5) = %d, want 10", got)
	}
}
