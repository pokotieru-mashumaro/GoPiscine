package piscine

func BasicAtoi(s string) int {
	count := 0
	for _, c := range s {
		count *= 10
		count += int(c - '0')
	}
	return count
}
