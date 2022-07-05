package validperfectsquare

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	input := []int{1, 2, 3, 4, 8}
	want := []bool{true, false, false, true, false}
	for i := 0; i < len(input); i++ {
		got := isPerfectSquare(input[i])
		if got != want[i] {
			t.Errorf("isPerfectSquare(%d) = %t, want %t", input[i], got, want[i])
		}
	}
}
