package spiralmatrix

func spiralOrder(matrix [][]int) []int {
	//matrix bound
	m := len(matrix)
	n := len(matrix[0])
	upper_bound := 0
	lower_bound := m - 1

	left_bound := 0
	right_bound := n - 1

	var res []int
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			for i := left_bound; i <= right_bound; i++ {
				res = append(res, matrix[upper_bound][i])
			}
			upper_bound++
		}
		if left_bound <= right_bound {
			for j := upper_bound; j <= lower_bound; j++ {
				res = append(res, matrix[j][right_bound])
			}
			right_bound--
		}
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			lower_bound--
		}
		if left_bound <= right_bound {
			for j := lower_bound; j >= upper_bound; j-- {
				res = append(res, matrix[j][left_bound])
			}
			left_bound++
		}

	}

	return res
}
