package main

import (
	"fmt"
)

func main() {
	// SHOW OMIT
	// 1 - 23
	// 2 - AA
	// 3 - 67
	code := 0x23AA67

	const (
		mask03, shift03 = 0xFF << (iota * 0x8), iota * 0x8
		mask02, shift02 = 0xFF << (iota * 0x8), iota * 0x8
		mask01, shift01 = 0xFF << (iota * 0x8), iota * 0x8
	)

	fmt.Printf(
		"%X %X %X\n",
		(code&mask01)>>shift01,
		(code&mask02)>>shift02,
		(code&mask03)>>shift03,
	)
	// END OMIT
}
