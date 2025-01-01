package piscine

func IsLow(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsLower(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsLow(c) {
			return false
		}
	}
	return true
}
