package stringtointegeratoi

import "math"

func myAtoi(s string) int {
	r := []byte(s)
	var intArray = make([]byte, len(r))
	k := 0
	other := 0
	sign := 1
	haveSign := 0

	for i := 0; i < len(r); i++ {
		if r[i] == ' ' && k == 0 && other > 0 {

			break
		}
		if r[i] == '.' {
			break
		}

		if r[i] == ' ' && haveSign > 0 {
			break
		}
		if r[i] == ' ' && k == 0 {
			continue
		}
		if r[i] != '-' && r[i] != '+' && (r[i] > '9' || r[i] < '0') {
			other++
		}
		if r[i] == '-' {
			if k > 0 {
				break
			}
			haveSign++
			sign = -1
		}
		if r[i] == '+' {
			if k > 0 {
				break
			}
			haveSign++
			sign = 1
		}
		if r[i] <= '9' && r[i] >= '0' {
			if other > 0 {
				break
			}
			intArray[k] = r[i]
			k++
		}
	}
	if haveSign > 1 {
		return 0
	}
	intArray2 := intArray[:k]
	if k > 0 {
		var result int
		for j := 0; j < len(intArray2); j++ {
			v := int(intArray[j]-'0') * sign
			result = result*10 + v

			if result > math.MaxInt32 {

				return math.MaxInt32
			}
			if result < -math.MaxInt32-1 {
				return -math.MaxInt32 - 1
			}
		}
		return result
	}
	return 0

}
