package piscine

func IsPrint(r rune) bool {
	if r >= ' ' && r <= '~' {
		return true
	}
	return false
}

func IsPrintable(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if !IsPrint(c) {
			return false
		}
	}
	return true
}
