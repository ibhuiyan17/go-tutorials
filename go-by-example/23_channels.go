package main

import (
	"fmt"
)


/*
	goroutine sends a ping message to main routine through a channel
*/
func main() {
	messages := make(chan string) // a channel
	
	go func() {
		messages <- "ping" // send to channel (blocking by defualt)
	}()

	msg := <-messages // recv from channel (also blocking by default)
	fmt.Println("rcved msg:", msg)

}