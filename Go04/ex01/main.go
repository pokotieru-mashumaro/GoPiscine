package main

import (
	"ex01/test/ft"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		for _, c := range os.Args[i] {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
