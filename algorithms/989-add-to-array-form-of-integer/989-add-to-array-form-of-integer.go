package addtoarrayformofinteger

func addToArrayForm(num []int, k int) []int {

	overFlow := 0
	remain := k % 10
	for i := len(num) - 1; i >= 0; i-- {
		v := num[i] + overFlow + remain
		if v >= 10 {
			overFlow = 1
			num[i] = v % 10
		} else {
			overFlow = 0
			num[i] = v
		}
		k = k / 10
		remain = k % 10

	}
	for k > 0 {
		v := k%10 + overFlow
		flow := []int{v % 10}
		num = append(flow, num...)

		if v >= 10 {
			overFlow = 1
		} else {
			overFlow = 0
		}

		k = k / 10

	}
	if overFlow == 1 {
		flow := []int{1}
		return append(flow, num...)
	}

	return num
}
