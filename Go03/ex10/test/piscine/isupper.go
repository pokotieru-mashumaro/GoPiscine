package piscine

func IsUpp(r rune) bool {
	if r >= 'A' && r <= 'Z' {
		return true
	}
	return false
}

func IsUpper(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsUpp(c) {
			return false
		}
	}
	return true
}
