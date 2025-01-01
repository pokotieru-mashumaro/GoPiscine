package piscine

func IsLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func ToUpper(s string) string {
	ret := ""
	for _, c := range s {
		if IsLower(c) {
			ret += string(c - ('a' - 'A'))
		} else {
			ret += string(c)
		}
	}
	return ret
}
