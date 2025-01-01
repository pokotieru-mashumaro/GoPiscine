package piscine

import "ex16/test/ft"

func IsDuplicate(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := i + 1; j < len(str); j++ {
			if str[i] == str[j] {
				return true
			}
		}
	}
	return false
}

func Calc(nbr int, base string) {
	if nbr <= 0 {
		return
	}
	digits := len(base)
	Calc(nbr/digits, base)
	ft.PrintRune(rune(base[nbr%digits]))
}

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 || IsDuplicate(base) {
		ft.PrintRune('F')
		ft.PrintRune('N')
		return
	}
	if nbr < 0 {
		nbr *= -1
		ft.PrintRune('-')
	}
	Calc(nbr, base)
}
