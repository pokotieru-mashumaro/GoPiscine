package main

import (
	"ex07/test/piscine"
	"fmt"
)

func main() {
	s := "Hello World!"
	s = piscine.StrRev(s)
	fmt.Println(s)
}
