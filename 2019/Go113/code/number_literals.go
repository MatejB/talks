package main

import (
	"fmt"
)

const (
	binary = 0b_111_111_111
	octal = 0o_7_7_7
)

func main() {
	fmt.Printf("%o\n", binary)
	fmt.Printf("%o\n", octal)
}
