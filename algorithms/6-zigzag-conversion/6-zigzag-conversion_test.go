package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	input := []string{
		"ABC",
		"AB",
		"PAYPALISHIRING",
		"PAYPALISHIRING",
		"A",
	}
	input1 := []int{
		1,
		1,
		3,
		4,
		1,
	}
	want := []string{
		"ABC",
		"AB",
		"PAHNAPLSIIGYIR",
		"PINALSIGYAHRPI",
		"A",
	}
	for i, v := range input {
		if got := convert(v, input1[i]); got != want[i] {
			t.Errorf("convert(%v, %v) = %v; want %v", v, input1[i], got, want[i])
		}
	}
}
