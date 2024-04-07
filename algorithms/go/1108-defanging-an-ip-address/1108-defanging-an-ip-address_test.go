package defanginganipaddress

import "testing"

func TestDefangIPaddr(t *testing.T) {
	input := "1.1.1.1"
	want := "1[.]1[.]1[.]1"
	got := defangIPaddr(input)
	if got != want {
		t.Fatalf(`defangIPaddr function failed, want %s, got %s .`, want, got)
	}
}
