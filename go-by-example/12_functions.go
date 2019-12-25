package main

import "fmt"

// format func name(params) return-type

func addTwo(a int, b int) int {
	return a + b
}

func addThree(a, b, c int) int {
	return addTwo(addTwo(a, b), c)
}

func main() {
	fmt.Println("1+2 =", addTwo(1, 2))
	fmt.Println("1+2+3 =", addThree(1,2,3))
}
