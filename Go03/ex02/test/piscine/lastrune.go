package piscine

func LastRune(s string) rune {
	len := len(s) - 1
	if len == 0 {
		return 0
	}
	rs := []rune(s)
	return rs[len]
}
