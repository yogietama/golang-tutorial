package main

import "fmt"

func goroutine(idx int) {
	fmt.Println("from go routine", idx)
}

func main() {
	// go routine is like thread

	for i := 0; i < 5; i++ {
		go goroutine(i)
	}

	fmt.Println("from main()")
}
