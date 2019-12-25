package main

import "fmt"

// intSeq is a function that returns a function that returns an int
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()	// intializes nextInt w/ i = 0 (closure of i)
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	another := intSeq()	// new instance
	fmt.Println(another())

}
