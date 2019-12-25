package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)

	scanner := bufio.NewScanner(resp.Body) // loads response body into scanner struct

	for scanner.Scan() { // continue until EOF
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil { // return if any other error
		panic(err)
	}
}
