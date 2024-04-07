package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	input := "()"
	want := true
	got := isValid(input)
	if got != want {
		t.Fatalf(`isValid function failed, want %v, got %v`, want, got)
	}
}
