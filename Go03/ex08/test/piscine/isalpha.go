package piscine

func IsLower(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsAlphaC(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func IsAlpha(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsAlphaC(c) {
			return false
		}
	}
	return true
}
