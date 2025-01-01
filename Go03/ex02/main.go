package main

import (
	"ex02/test/ft"
	"ex02/test/piscine"
)

func main() {
	ft.PrintRune(piscine.LastRune("Hello!"))
	ft.PrintRune(piscine.LastRune("Salut!"))
	ft.PrintRune(piscine.LastRune("Ola!"))
	ft.PrintRune('\n')
}
