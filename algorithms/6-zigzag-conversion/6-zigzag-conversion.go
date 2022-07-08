package zigzagconversion

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	zigs := make([][]string, numRows)
	for i := range zigs {
		zigs[i] = make([]string, len(s))

	}
	row, col := 0, 0
	for i := 0; i < len(s); i++ {
		if col < numRows {
			zigs[col][row] = string(s[i])
			col++
			continue
		} else if col == numRows {
			col--
			for j := 0; j < numRows-1; j++ {
				col--
				row++
				zigs[col][row] = string(s[i])
				i++
				if i == len(s) {
					break
				}
			}
			i--
			col++
		}
	}
	var res = ""
	for m := 0; m < numRows; m++ {
		for n := 0; n < len(s); n++ {
			v := zigs[m][n]
			if v != "" {
				res += v
			}
		}
	}
	return res
}
