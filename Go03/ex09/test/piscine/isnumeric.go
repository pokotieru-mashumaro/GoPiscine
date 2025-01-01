package piscine

func IsNum(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}
	return false
}

func IsNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsNum(c) {
			return false
		}
	}
	return true
}
