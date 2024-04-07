package sort

import "testing"

func TestFrequencySort(t *testing.T) {
	input := []string{
		"loveleetcode",
		"cccaaa",
		"",
	}
	want := []string{
		"eeeeoollvtdc",
		"cccaaa",
		"",
	}
	for i, v := range input {
		got := frequencySort(v)
		if got != want[i] {
			t.Fatalf(`frequencySort function failed, want %s, got %s`, want[i], got)
		}
	}
}
