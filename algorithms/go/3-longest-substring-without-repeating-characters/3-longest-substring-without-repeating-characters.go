package longestsubstringwithoutrepeatingcharacters

//two example
//all the not same, all the same
func lengthOfLongestSubstring(s string) int {
	var window = make(map[byte]int)
	//left and right window size
	left, right := 0, 0
	res := 0
	for right < len(s) {
		r := s[right]
		right++
		window[r]++
		//why for or not if,for it is a sliding window, go for next
		for window[r] > 1 {
			l := s[left]
			left++
			window[l]--
		}
		//compare now window size and history window size
		if res < right-left {
			res = right - left
		}
	}

	return res
}
