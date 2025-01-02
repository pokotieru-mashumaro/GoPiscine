package piscine

import "ex04/test/ft"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		for _, c := range a[i] {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
