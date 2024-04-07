package implementstrstr

import "testing"

func TestStrStr(t *testing.T) {
	input := []string{
		"hello",
		"mississippi",
	}

	input2 := []string{
		"ll",
		"issip",
	}
	want := []int{
		2,
		4,
	}

	for i := range input {
		got := strStr(input[i], input2[i])
		if got != want[i] {
			t.Fatalf(`strStr function failed, want %d, got %d.`, want[i], got)
		}
	}
}
