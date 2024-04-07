package simplifypath

import "testing"

func TestSimplifyPath(t *testing.T) {
	input := []string{
		"/home/",
		"/home//foo/",
	}
	want := []string{
		"/home",
		"/home/foo",
	}
	for i := range input {
		got := simplifyPath(input[i])
		if got != want[i] {
			t.Fatalf(`simplifyPath function failed, want %s, got %s.`, want[i], got)
		}
	}
}
