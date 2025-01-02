package piscine

import (
	"ex02/test/ft"
	"os"
)

func Display(s string) {
	for i := 0; i < len(s); i++ {
		ft.PrintRune(rune(s[i]))
	}
	ft.PrintRune('\n')
}

func DisplayFile(filepath string) {
	fp, err := os.Open(filepath)
	if err != nil {
		Display("File is error")
		return
	}

	buf := make([]byte, 1)
	tmp := ""
	for {
		n, err := fp.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			Display("Read error")
			return
		}
		tmp += string(buf)
	}
	fp.Close()
	Display(tmp)
}
