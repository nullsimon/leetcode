package dividetwointegers

func divide(dividend int, divisor int) int {
	neg := false
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		neg = true
	}

	dividend = Abs(dividend)
	divisor = Abs(divisor)

	dividend = dividend - divisor
	res := 1
	for isContinue(dividend) {
		dividend = dividend - divisor
		res++
	}

	if dividend != 0 {
		res -= 1
	}

	if neg {
		res = -res
	}

	if res > 1<<31-1 {
		return 1<<31 - 1
	} else {
		return res
	}

}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func isContinue(i int) bool {
	return i >= 1
}
