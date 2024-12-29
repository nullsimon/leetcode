package validpalindrome

func isPalindrome(s string) bool {
	list := []rune{}
	for _, r := range s {
		if r <= 'Z' && r >= 'A' {
			list = append(list, r+32)
		} else if r <= 'z' && r >= 'a' {
			list = append(list, r)
		} else if r <= '9' && r >= '0' {
			list = append(list, r)
		}

	}
	for i := len(list) - 1; i > i/2; i-- {
		if list[i-1] != list[len(list)-i] {
			return false
		}
	}
	return true

}
