package piscine

func Atoi(s string) int {
	num := 0
	sign := 1
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -1
			continue
		} else if i == 0 && c == '+' {
			continue
		} else if c < '0' || c > '9' {
			return 0
		}
		num *= 10
		num += int(c - '0')
	}
	return num * sign
}
