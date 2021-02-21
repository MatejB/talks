package main

import (
	"fmt"
	"time"
)

func main() {
	// SHOW OMIT
	const millisecond = time.Second / 1e3
	fmt.Println(millisecond)

	bigBuffer := make([]byte, 512+1e6)
	fmt.Println(len(bigBuffer))
	// END OMIT
}
