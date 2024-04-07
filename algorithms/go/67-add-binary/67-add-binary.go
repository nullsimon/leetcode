package addbinary

import "fmt"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	var ch = make(chan string, 10000)
	overflow := 0
	for i >= 0 || j >= 0 {
		if i < 0 && j >= 0 {
			num := int(b[j]-'0') + overflow
			if num == 2 {
				overflow = 1
				ch <- "0"
			} else {
				overflow = 0

				ch <- fmt.Sprint(num)
			}
			j--

		} else if j < 0 && i >= 0 {
			num := int(a[i]-'0') + overflow
			if num == 2 {
				overflow = 1
				ch <- "0"
			} else {
				overflow = 0

				ch <- fmt.Sprint(num)
			}
			i--
		} else {
			num := int(a[i]-'0') + int(b[j]-'0') + overflow
			if num == 2 {
				overflow = 1
				ch <- "0"
			} else if num == 3 {
				overflow = 1
				ch <- "1"
			} else {
				overflow = 0

				ch <- fmt.Sprint(num)
			}

			i--
			j--
		}
	}
	if overflow == 1 {
		ch <- "1"
	}
	var res = ""
	for len(ch) > 0 {
		cur := <-ch
		fmt.Println(cur)
		res = cur + res
	}

	return res
}
