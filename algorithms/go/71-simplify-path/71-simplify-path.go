package simplifypath

import "strings"

func simplifyPath(path string) string {
	split := strings.Split(path, "/")
	if len(path) == 1 {
		return path
	}
	var stack []string

	for j := 0; j < len(split); j++ {
		var text = split[j]

		if text == "" {
			continue
		} else if text == "." {
			continue
		} else if text == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		} else {
			stack = append(stack, text)
		}
	}
	var res = ""

	for i := 0; i < len(stack); i++ {

		res = res + "/" + stack[i]
	}
	if len(stack) == 0 {
		return "/"
	}
	return res

}
