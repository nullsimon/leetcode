package integertoroman

/*
* intToRoman Tips
* 1 using int slice to ordered map
* 2 there are different nunmber, 4 in a row should go up, 9 should go up--should ignore 5
*
 */
func intToRoman(num int) string {
	var intRomans = map[int]string{
		1000: "M",
		500:  "D",
		100:  "C",
		50:   "L",
		10:   "X",
		5:    "V",
		1:    "I",
	}
	var ints = []int{1000, 500, 100, 50, 10, 5, 1}

	res := ""

	for num > 0 {
		for k, v := range ints {
			j := num / v
			left := num % v
			num = left
			if j == 4 {
				res += intRomans[v]
				res += intRomans[ints[k-1]]
			} else {
				for ; j > 0; j-- {
					res = res + intRomans[v]
				}
			}
			five := false
			fiveNumber := v
			for fiveNumber > 5 {
				fiveNumber = fiveNumber / 10
			}
			if fiveNumber == 5 {
				five = true
			}
			if left*10/v >= 9 && !five {
				res += intRomans[ints[k+2]]
				res += intRomans[v]
				num = num - v + ints[k+2]
			}
		}
	}
	return res
}
