package integertoroman

import (
	"testing"
)

func TestIntToRoman(t *testing.T) {
	input := []int{49, 45, 3, 1, 4, 7, 1994}
	want := []string{"XLIX", "XLV", "III", "I", "IV", "VII", "MCMXCIV"}
	for i := 0; i < len(input); i++ {
		//got := intToRoman(input[i])
		got := intToRomanHardCode(input[i])
		if got != want[i] {
			t.Fatalf(`integer to roman func failed: want %s, got %s`, want[i], got)
		}
	}
}
