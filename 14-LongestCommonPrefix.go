package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"vim-go", "vim"}))
}

func longestCommonPrefix(strs []string) string {

	common := ""
	if len(strs) < 1 {
		return common
	}
	if len(strs) < 2 {
		return strs[0]
	}
	if len(strs[0]) < 1 {
		return common
	}

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) {
				return common
			}
			if strs[0][i] != strs[j][i] {
				return common
			}
		}
		common = common + string(strs[0][i])
	}
	return common
}
