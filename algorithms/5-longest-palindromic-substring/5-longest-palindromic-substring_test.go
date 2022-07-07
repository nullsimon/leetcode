package longestpalindromicsubstring

import "testing"

func TestLongestPalindrome(t *testing.T) {
	input := []string{
		"babad",
		"cbbd",
		"a",
	}
	want := []string{
		"bab",
		"bb",
		"a",
	}
	for i, v := range input {
		if got := longestPalindrome(v); got != want[i] {
			t.Errorf("longestPalindrome(%q) = %q; want %q", v, got, want[i])
		}
	}
}
