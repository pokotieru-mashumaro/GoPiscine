package main

import (
	"ex00/test/ft"
	"ex00/test/piscine"
)

func main() {
	ft.PrintRune(piscine.FirstRune("Hello!"))
	ft.PrintRune(piscine.FirstRune("Salut!"))
	ft.PrintRune(piscine.FirstRune("Ola!"))
	ft.PrintRune('\n')
}
