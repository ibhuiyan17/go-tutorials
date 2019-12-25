package main

import "fmt"

func main() {
	s := make([]int, 0)
	fmt.Println("empty:", s)

	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
	fmt.Println("after:", s)
	fmt.Println("len:", len(s))

	sl := s[1:5]
	fmt.Println("slice of s [1,5):", sl)
}
