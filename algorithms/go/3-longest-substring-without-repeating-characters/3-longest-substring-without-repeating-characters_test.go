package longestsubstringwithoutrepeatingcharacters

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	input := []string{
		"abcde",
		"abcabcbb",
		"bbbbb",
		"pwwkew",
	}
	want := []int{
		5,
		3,
		1,
		3,
	}
	for i := range input {
		got := lengthOfLongestSubstring(input[i])
		if got != want[i] {
			t.Fatalf(`lengthOfLongestSubstring function failed, want %d, got %d.`, want[i], got)
		}
	}
}
