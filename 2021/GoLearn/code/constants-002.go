package main

import "fmt"

func main() {
	// SHOW OMIT
	fmt.Printf("%T %v\n", "", "")
	fmt.Printf("%T %v\n", 0, 0)
	fmt.Printf("%T %v\n", 0.0, 0.0)
	fmt.Printf("%T %v\n", true, true)
	fmt.Printf("%T %v\n", 'a', 'a')
	fmt.Printf("%T %v\n", 0i, 0i)
	// END OMIT
}
