package main

import "fmt"

func main() {
	fmt.Println("sum")
	nums := []int{2,4,6,8}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	fmt.Println("using index")
	for i, num := range nums {
		fmt.Println(i, num)
	}

	fmt.Println("iterating over map")
	m := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range m {
		fmt.Printf("%s -> %s\n", k, v)
	}

	fmt.Println("iterating over string 'hello'")
	for _, c := range "hello" {
		fmt.Println(c)
	}
}
