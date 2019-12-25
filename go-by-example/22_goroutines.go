package main

import (
	"fmt"
	"time"
)

func f(caller string) {
	for i := 0; i < 20; i++ {
		fmt.Println(caller, ":", i)
	}
}

func main() {

	// separate routines
	go f("goroutine 1") 
	go f("goroutine 2")
	
	// executing in the main routine
	f("direct")
		
	// anonymous func. goroutine
    go func(msg string) {
        fmt.Println(msg)
    }("going")

	time.Sleep(time.Second)
	time.Sleep(time.Second)
	fmt.Println("done")
}