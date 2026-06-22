package notviable

import "testing"

func TestConcat(t *testing.T) {
	if got := Concat("go", "lang"); got != "golang" {
		t.Fatalf("Concat(\"go\", \"lang\") = %q, want \"golang\"", got)
	}
}
