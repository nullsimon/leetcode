package timeneededtoinformallemployees

import "testing"

func TestNumOfMinutes(t *testing.T) {
	/*7
	6
	[1,2,3,4,5,6,-1]
	[0,6,5,4,3,2,1]
	*/
	// input1 := 1
	// input2 := 0
	// input3 := []int{-1}
	// input4 := []int{0}

	input1 := 7
	input2 := 6
	input3 := []int{1, 2, 3, 4, 5, 6, -1}
	input4 := []int{0, 6, 5, 4, 3, 2, 1}
	want := 21

	got := numOfMinutes(input1, input2, input3, input4)
	if got != want {
		t.Fatalf(`roman to integer func failed: want %d, got %d`, want, got)
	}

}
