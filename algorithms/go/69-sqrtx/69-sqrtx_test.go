package sqrtx

import "testing"

func TestMySqrt(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 8}
	want := []int{0, 1, 1, 1, 2, 2}
	for i := 0; i < len(input); i++ {
		got := mySqrt(input[i])
		if got != want[i] {
			t.Errorf("mySqrt(%d) = %d, want %d", input[i], got, want[i])
		}
	}
}
