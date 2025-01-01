package main

import (
	"ex12/test/piscine"
	"fmt"
)

func main() {
	fmt.Println(piscine.IsPrintable("Hello"))
	fmt.Println(piscine.IsPrintable("Hello\n"))
}
