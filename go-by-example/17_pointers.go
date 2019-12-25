package main

import "fmt"

func valFunc(val int) { // makes a copy
	val = 0
}

func ptrFunc(ptr *int) { // will change calling val
	*ptr = 0
}

func main() {
	i := 1
	valFunc(i) // wont do anything to i

	fmt.Println(i)

	ptrFunc(&i) // affects i
	fmt.Println(i)
}
