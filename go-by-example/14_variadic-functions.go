package main

import "fmt"


// takes in a list of arguments
func sumSlice(nums ...int) int {
	fmt.Println("summing", nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	fmt.Println(sumSlice(2,5,1,3,7,8,9)) // arbitrary number of args

	nums := []int{3,6,1,2}
	fmt.Println(sumSlice(nums...))	// slice as arg
}
