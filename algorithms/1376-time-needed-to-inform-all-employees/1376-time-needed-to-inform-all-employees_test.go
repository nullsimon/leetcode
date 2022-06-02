package timeneededtoinformallemployees

import "testing"

func TestNumOfMinutes(t *testing.T) {
	input1 := 1
	input2 := 0
	input3 := []int{-1}
	input4 := []int{0}
	want := 0

	got := numOfMinutes(input1, input2, input3, input4)
	if got != want {
		t.Fatalf(`roman to integer func failed: want %d, got %d`, want, got)
	}

}
