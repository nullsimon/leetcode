package thetimewhenthenetworkbecomesidle

import "testing"

func TestNetworkBecomesIdle(t *testing.T) {
	input := [][]int{{0, 1}, {0, 2}, {1, 2}}
	input2 := []int{0, 10, 10}
	want := 3

	got := networkBecomesIdle(input, input2)
	if got != want {
		t.Fatalf(`roman to integer func failed: want %d, got %d`, want, got)
	}

}
