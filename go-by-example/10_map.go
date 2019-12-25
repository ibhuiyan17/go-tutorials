package main

import "fmt"

func main() {
	// delcared with make(map[key-type]value-type)
	hashmap := make(map[string]int)

	hashmap["k1"] = 5
	hashmap["k2"] = 8
	fmt.Println("map:", hashmap)
	fmt.Println("map len:", len(hashmap))

	v1 := hashmap["k1"]
	fmt.Println("v1:", v1)

	hashmap["k3"] = 21

	_, prs := hashmap["k3"]		// check if present
	fmt.Println("prs:", prs)
	delete(hashmap, "k3")
	_, prs2 := hashmap["k3"]	// check if present
	fmt.Println("prs2", prs2)

}
