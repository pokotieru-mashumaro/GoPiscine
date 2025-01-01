package piscine

func NRune(s string, n int) rune {
	if n < 0 || n > len(s) {
		return 0
	}
	rs := []rune(s)
	return rs[n-1]
}
