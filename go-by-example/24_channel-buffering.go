package main

import (
	"fmt"
)

/*
	By default, unbuffered channel only accepts if there is
	a corresponding receive. Here the receives are not concurrent
	but allows 3 values to be buffered.
*/
func main() {
	channel := make(chan string, 3) // can buffer up to 3 strings

	channel <- "this"
	channel <- "is"
	channel <- "buffered"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

}