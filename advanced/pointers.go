package main

import "fmt"

func main() {
	var p *string
	fmt.Println(p == nil)

	name := "nakama"
	// p = name --> error
	p = &name

	fmt.Println(p)  // some var address
	fmt.Println(&p) // some var address
	fmt.Println(*p)

	// &p = "someval" error
	*p = "updateVal"
	fmt.Println(*p)

}
