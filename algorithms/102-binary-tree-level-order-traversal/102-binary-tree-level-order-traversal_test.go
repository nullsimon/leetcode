package binarytreelevelordertraversal

import (
	"testing"
)

func TestLevelOrder(t *testing.T) {
	input := []int{49, 45, 3, 1, 4, 7, 1994}
	want := []string{"XLIX", "XLV", "III", "I", "IV", "VII", "MCMXCIV"}
	for i := 0; i < len(input); i++ {
		//got := intToRoman(input[i])
		// got := levelOrder(input[i])
		// if got != want[i] {
		// 	t.Fatalf(`integer to roman func failed: want %s, got %s`, want[i], got)
		// }
		t.Logf(`this is what i want %s`, want[i])
		var leverTreeSlice [][]int
		var slice []int
		slice = append(slice, 1)
		leverTreeSlice = append(leverTreeSlice, slice)
		t.Log(leverTreeSlice)
	}
}
