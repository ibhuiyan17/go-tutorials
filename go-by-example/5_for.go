package main

import "fmt"

func main() {
	// only for loops exist in Go; 3 types

	// while-style for-loop
	fmt.Println("while-style")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// classic for-loop w/ conditions
	fmt.Println("classic w/ conds")
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

	// infinite while(true) style
	fmt.Println("infinite")
	i = 0
	for {
		fmt.Println("loops infinitely until broken")
		break
	}



	// random, find odds between [0, 10)
	fmt.Println("odds")
	for num := 0; num < 10; num++ {
		if num % 2 == 1 {
			fmt.Println(num)
		}
	}

}
