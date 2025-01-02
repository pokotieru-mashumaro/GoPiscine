package main

import (
	"ex00/test/ft"
	"os"
)

func main() {
	execPath := os.Args[0]
	execName := execPath
	for i := len(execPath) - 1; i >= 0; i-- {
		if execPath[i] == '/' {
			execName = execPath[i+1:]
			break
		}
	}
	for _, c := range execName {
		ft.PrintRune(c)
	}
}
