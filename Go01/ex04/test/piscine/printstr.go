package piscine

import "ex04/test/ft"

func PrintStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}
