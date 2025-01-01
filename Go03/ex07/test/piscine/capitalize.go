package piscine

func IsLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsAlNum(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	} else if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	ret := ""
	isFirst := true
	for _, c := range s {
		if !IsAlNum(c) {
			ret += string(c)
			isFirst = true
		} else {
			if isFirst && IsLower(c) {
				ret += string(c - ('a' - 'A'))
			} else {
				ret += string(c)
			}
			isFirst = false
		}
	}
	return ret
}
