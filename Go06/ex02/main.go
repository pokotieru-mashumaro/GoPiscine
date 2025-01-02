package main

import (
	"ex02/test/piscine"
	"os"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		piscine.Display("File name missing")
		return
	} else if len(args) > 2 {
		piscine.Display("Too many arguments")
		return
	}
	piscine.DisplayFile(args[1])
}
