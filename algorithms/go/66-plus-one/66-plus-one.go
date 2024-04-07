package plusone

func plusOne(digits []int) []int {
	overFlow := 1
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i] + overFlow
		if v == 10 {
			overFlow = 1
			digits[i] = 0
		} else {
			overFlow = 0
			digits[i] = v
		}
	}
	if overFlow == 1 {
		flow := []int{1}
		return append(flow, digits...)
	}
	return digits
}
