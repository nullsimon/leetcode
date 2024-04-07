package rotateimage

func rotate(matrix [][]int) [][]int {
	length := len(matrix)

	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			//swap matrix[i][j] matrix[j][i]
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	for k := 0; k < length; k++ {
		row := matrix[k]
		i, j := 0, len(row)-1
		for i < j {
			tmp := row[i]
			row[i] = row[j]
			row[j] = tmp
			i++
			j--
		}
	}
	return matrix
}
