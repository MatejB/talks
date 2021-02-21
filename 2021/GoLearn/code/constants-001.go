package main

import "fmt"

func main() {
	// SHOW OMIT
	const foo = "foo"        // untyped string constant
	const bar string = "bar" // typed string constant

	type myString string

	var f myString = foo
	var b myString = bar

	fmt.Println(f)
	fmt.Println(b)
	// END OMIT
}
