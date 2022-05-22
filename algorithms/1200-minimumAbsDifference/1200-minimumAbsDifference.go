package minimumabsdifference

import "sort"

func minimumAbsDifference(arr []int) [][]int {

	if len(arr) < 2 {
		return nil
	}
	sort.Ints(arr)

	minimum := getTheMinimum(arr)
	pairs := make([][]int, 0, 16)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == minimum {
			pair := []int{arr[i], arr[i+1]}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func getTheMinimum(arr []int) int {
	minimum := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] < minimum {
			minimum = arr[i+1] - arr[i]
		}
	}
	return minimum
}
