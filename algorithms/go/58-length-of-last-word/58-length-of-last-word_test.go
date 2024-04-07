package lengthOfLastWord

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	input := []string{
		"hello world",
		"hello",
		"wwww wwww www    ",
		"",
		" ",
	}
	want := []int{
		5,
		5,
		3,
		0,
		0,
	}
	for i, v := range input {
		got := lengthOfLastWord(v)
		if got != want[i] {
			t.Fatalf(`lengthOfLastWord function failed, want %T, got %T`, want[i], got)
		}
	}
}
