package piscine

import "ex00/test/ft"

func PrintAlphabet() {
	for c := 'a'; c <= 'z'; c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
