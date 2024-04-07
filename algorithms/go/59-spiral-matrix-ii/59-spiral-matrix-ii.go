package spiralmatrixii

func generateMatrix(n int) [][]int {
	upper_bound, left_bound := 0, 0
	lower_bound, right_bound := n-1, n-1

	number := 0
	var matrix = make([][]int, n)
	for k := range matrix {
		matrix[k] = make([]int, n)
	}
	for number < n*n {
		if upper_bound <= lower_bound {
			for j := left_bound; j <= right_bound; j++ {
				number++
				matrix[upper_bound][j] = number
			}
			upper_bound++
		}
		if left_bound <= right_bound {
			for j := upper_bound; j <= lower_bound; j++ {
				number++
				matrix[j][right_bound] = number
			}
			right_bound--
		}
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				number++
				matrix[lower_bound][j] = number
			}
			lower_bound--
		}
		if left_bound <= right_bound {
			for j := lower_bound; j >= upper_bound; j-- {
				number++
				matrix[j][left_bound] = number
			}
			left_bound++
		}

	}
	return matrix

}
