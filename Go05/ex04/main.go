package main

import (
	"ex04/test/piscine"
)

func main() {
	a := piscine.SplitWhiteSpaces("Hello how are you?")
	piscine.PrintWordsTables(a)
}
