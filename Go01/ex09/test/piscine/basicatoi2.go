package piscine

func BasicAtoi2(s string) int {
	count := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		count *= 10
		count += int(c - '0')
	}
	return count
}
