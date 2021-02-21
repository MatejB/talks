package main

import (
	"fmt"
)

func main() {
	// SHOW OMIT
	const (
		zero = iota // 0
		one         // 1
		two         // 2
	)
	fmt.Println(zero, one, two)
	// END OMIT
}
