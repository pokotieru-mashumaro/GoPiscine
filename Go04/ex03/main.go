package main

import (
	"ex03/test/ft"
	"ex03/test/piscine"
	"os"
)

func main() {
	sortArgs := piscine.SortParams(os.Args)
	for i := 1; i < len(sortArgs); i++ {
		for _, c := range sortArgs[i] {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
