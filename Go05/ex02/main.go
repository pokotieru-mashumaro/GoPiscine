package main

import (
	"ex02/test/piscine"
	"fmt"
)

func main() {
	test := []string{"Hello", "how", "are", "you?"}
	fmt.Println(piscine.ConcatParams(test))
}
