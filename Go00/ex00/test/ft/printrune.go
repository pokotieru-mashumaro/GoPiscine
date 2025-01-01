package ft

import "os"

func PrintRune(r rune) {
	os.Stdout.Write([]byte(string(r)))
}
