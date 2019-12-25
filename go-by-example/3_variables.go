package main

import "fmt"

func main() {
	// vars can be explicitly or implicitly declared in GO

	var a = "initial"	// implicit
	fmt.Println(a)

	var b, c int = 1, 2	// explicit
	fmt.Println(b, c)

	var d = true		// implicit
	var e bool = false	// explicit
	fmt.Println(d, e)

	f := "apple"		// declare & initialize
	fmt.Println(f)

	g, h := "initialized", true	// multiple declare and initialize
	fmt.Println(g, h)
}
