package firstuniquecharacterinastring

import "testing"

func TestFirstUniqChar(t *testing.T) {
	input := []string{
		"leetcode",
		"loveleetcode",
		"aabb",
		"a",
	}
	want := []int{
		0,
		2,
		-1,
		0,
	}

	for i, v := range input {
		got := firstUniqChar(v)
		if got != want[i] {
			t.Fatalf(`firstUniqChar function failed, want %d, got %d`, want[i], got)
		}
	}
}
