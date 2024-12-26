package validpalindrome

func isPalindrome(s string) bool {
	trimmed := ""
	for _, r := range s {
		if r <= 'Z' && r >= 'A' {
			trimmed = trimmed + string(r+32)
		} else if r <= 'z' && r >= 'a' {
			trimmed = trimmed + string(r)
		} else if r <= '9' && r >= '0' {
			trimmed = trimmed + string(r)
		}

	}
	reversed := ""
	list := []rune{}
	for _, r := range trimmed {
		list = append(list, r)
	}
	for i := len(list) - 1; i >= 0; i-- {
		reversed = reversed + string(list[i])
	}
	return trimmed == reversed
}
