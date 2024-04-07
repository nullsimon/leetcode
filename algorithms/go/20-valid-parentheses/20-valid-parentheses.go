package validparentheses

func isValid(s string) bool {

	if len(s)%2 != 0 {
		return false
	}

	var stack []string
	for _, i := range s {

		if i == '(' || i == '{' || i == '[' {
			//push
			stack = append(stack, string(i))

		} else {
			if len(stack) > 0 && stack[len(stack)-1] == rightOf(string(i)) {
				stack = stack[:len(stack)-1]
			} else {
				return false

			}
		}
	}
	return len(stack) == 0

}

func rightOf(i string) string {
	if i == ")" {
		return "("
	}
	if i == "}" {
		return "{"
	}
	return "["
}
