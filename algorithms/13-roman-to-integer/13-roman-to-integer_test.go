package romanToInteger

import "testing"

func TestRomanToInt(t *testing.T) {
	input := []string{"III", "I", "IV", "VII"}
	want := []int{3, 1, 4, 7}
	for i := 0; i < len(input); i++ {
		got := romanToInt(input[i])
		if got != want[i] {
			t.Fatalf(`roman to integer func failed: want %d, got %d`, want[i], got)
		}
	}
}
