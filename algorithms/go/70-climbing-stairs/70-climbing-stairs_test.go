package climbingstairs

import "testing"

func TestClimbStairs(t *testing.T) {
	input := []int{1, 2, 3, 45}
	expected := []int{1, 2, 3, 1836311903}
	for i := range input {
		got := climbStairs(input[i])
		if got != expected[i] {
			t.Errorf("climbStairs(%d) = %d, want %d", input[i], got, expected[i])
		}
	}
}
