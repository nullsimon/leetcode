package reverseinteger

import "testing"

func TestReverse(t *testing.T) {
	input := []int{123, 321, -223, 0}
	want := []int{321, 123, -322, 0}
	for i := 0; i < len(input); i++ {
		got := reverse(input[i])
		if got != want[i] {
			t.Fatalf(`integer reverse func failed: want %d, got %d`, want, got)
		}
	}
}
