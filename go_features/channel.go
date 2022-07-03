package main

import (
	"fmt"
	"time"
)

func insert(idx int, ch chan string) {
	fmt.Printf("Goroutine %d - Start Sending data \n", idx)
	ch <- fmt.Sprintf("Data %d", idx)
	fmt.Printf("Goroutine %d - End Sending data", idx)
}

func main() {
	// c := make(chan string) // unbuffer channle
	c := make(chan string, 2) // buffer channel
	defer close(c)

	for i := 0; i < 3; i++ {
		go insert(i, c)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Start Receiving data")

	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("End Receiving data")
}
