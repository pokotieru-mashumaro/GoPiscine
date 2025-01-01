package piscine

import "ex04/test/ft"

func PrintComb() {
	for left := 0; left <= 7; left++ {
		for middle := left + 1; middle <= 8; middle++ {
			for right := middle + 1; right <= 9; right++ {
				ft.PrintRune(rune(left) + '0')
				ft.PrintRune(rune(middle) + '0')
				ft.PrintRune(rune(right) + '0')
				if left != 7 {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
}
