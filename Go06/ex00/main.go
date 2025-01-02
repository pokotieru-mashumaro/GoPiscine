package main

import (
	"ex00/test/ft"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

func isEven(nbr int) bool {
	if even(nbr) == 1 {
		return true
	} else {
		return false
	}
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	lengthOfArg := len(os.Args) - 1
	if isEven(lengthOfArg) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
