package main

import "fmt"

func main() {
	/*
		defer func to delay func, in stack order
	*/

	defer fmt.Println("will be called last")

	fmt.Println("Will be called first")

	defer fmt.Println("will be called second last")

	fmt.Println("will be called second")

	if false {
		defer fmt.Println("will not be called")
	}
}
