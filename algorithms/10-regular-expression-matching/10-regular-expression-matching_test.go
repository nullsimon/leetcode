package regularexpressionmatching

import "testing"

func TestIsMatch(t *testing.T) {
	input := []string{
		"aab",
		"aa",
		"aa",
		"ab",
	}
	input1 := []string{
		"c*a*b",
		"a",
		"a*",
		".*",
	}
	want := []bool{
		true,
		false,
		true,
		true,
	}
	for i, v := range input {
		if got := isMatch(v, input1[i]); got != want[i] {
			t.Errorf("isMatch(%v, %v) = %v; want %v", v, input1[i], got, want[i])
		}
	}

}
