package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(work_channel chan int, num int) {
	fmt.Println("worker", num, "working...")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Println("worker", num, "done")

	work_channel <- num
}

func main() {
	num_workers := 5
	work_channel := make(chan int, num_workers)

	for worker_num := 0; worker_num < num_workers; worker_num++ {
		go worker(work_channel, worker_num)
	}

	fmt.Println(<-work_channel, "received")
	fmt.Println(<-work_channel, "received")
	fmt.Println(<-work_channel, "received")
	fmt.Println(<-work_channel, "received")
	fmt.Println(<-work_channel, "received")
}
