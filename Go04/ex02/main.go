package main

import (
	"ex02/test/ft"
	"os"
)

func main() {
	for i := len(os.Args) - 1; i >= 1; i-- {
		for _, c := range os.Args[i] {
			ft.PrintRune(c)
		}
		ft.PrintRune('\n')
	}
}
