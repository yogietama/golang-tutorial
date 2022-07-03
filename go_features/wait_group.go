package main

import (
	"fmt"
	"sync"
)

func goroutineWG(idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("from go routine", idx)
}

func main() {
	// go routine is like thread
	// use to wait until all goroutine finished
	var wg sync.WaitGroup
	wg.Add(5) // add according to total goroutine

	for i := 0; i < 5; i++ {
		go goroutineWG(i, &wg)
	}

	wg.Wait()
	fmt.Println("after all goroutine finished")
}
