package main

import (
	"ex06/test/piscine"
	"fmt"
)

func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", piscine.Split(s, "HA"))
}
