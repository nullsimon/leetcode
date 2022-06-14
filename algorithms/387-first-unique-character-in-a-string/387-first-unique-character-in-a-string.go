package firstuniquecharacterinastring

func firstUniqChar(s string) int {

	var freq = make([]int, 128)
	var ch = make(chan int, 100000)

	for i := 0; i < len(s); i++ {

		freq[s[i]]++
		if freq[s[i]] == 1 {
			ch <- i
		}
	}
	for len(ch) > 0 {
		cur := <-ch
		if freq[s[cur]] == 1 {
			return cur
		}
	}
	return -1

}
