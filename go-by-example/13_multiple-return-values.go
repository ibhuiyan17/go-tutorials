package main

import "fmt"

func addAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a, b := addAndMultiply(2,7)
	fmt.Println("2+7 =", a)
	fmt.Println("2*7 =", b)

	_, c := addAndMultiply(4,2)
	fmt.Println("4*2 =", c)

}
