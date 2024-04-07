package lengthOfLastWord

func lengthOfLastWord(s string) int {
	length := 0
	ignore := true
	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == ' ' {
			if ignore {
				continue
			} else {
				break
			}
		}
		ignore = false
		length++
	}
	return length
}
