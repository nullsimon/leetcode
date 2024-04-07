package regularexpressionmatching

func isMatch(s string, p string) bool {
	i, j := 0, 0
	for i < len(s)-1 && j < len(p)-1 {
		if s[i] == p[j] || p[j] == '.' {

			// if next is '*'
			if j+1 < len(p) && p[j+1] == '*' {
				if p[j] == '.' {
					j++
					i++
				} else {
					j++
					sv := s[i]
					for i+1 < len(s) && s[i+1] == sv {
						i++
					}
					j++
				}

			} else {
				i++
				j++
			}

		} else {
			if j+1 < len(p) && p[j+1] == '*' {
				j++
			} else {
				if p[j] == '*' {
					j++
					continue
				}
				return false
			}
		}
	}

	return len(s)-i == len(p)-j
}
