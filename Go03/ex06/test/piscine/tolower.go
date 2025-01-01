package piscine

func IsUpper(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func ToLower(s string) string {
	ret := ""
	for _, c := range s {
		if IsUpper(c) {
			ret += string(c + ('a' - 'A'))
		} else {
			ret += string(c)
		}
	}
	return ret
}
