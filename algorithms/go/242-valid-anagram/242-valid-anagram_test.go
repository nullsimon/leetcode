package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	input1 := []string{
		"anagram",
		"rat",
	}
	input2 := []string{
		"nagaram",
		"car",
	}
	want := []bool{
		true,
		false,
	}
	for i := range input1 {
		got := isAnagram(input1[i], input2[i])
		if got != want[i] {
			t.Fatalf(`isAnagram function failed, want %v, got %v.`, want, got)
		}
	}
}
