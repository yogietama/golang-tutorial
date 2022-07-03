package main

import "fmt"

func setFoo() error {
	return fmt.Errorf("some error")
}

func main() {
	x := 15
	var1 := "alpha"

	// if else condition
	if x == 10 {
		fmt.Println("x is 10")
	} else if x == 20 {
		fmt.Println("x is 20")
	} else {
		fmt.Println("other x value")
	}

	// direct if else
	if err := setFoo(); err != nil {
		fmt.Println("SetFoo func returns error")
	}

	// switch case
	switch var1 {
	case "aplha":
		fmt.Println("var1 == alpha")
	case "beta":
		fmt.Println("var1 == beta")
	default:
		fmt.Println("other var1 value")
	}

	// for loop with index
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// loop direct with data
	data := []string{"this", "is", "golang", "tutorial"}

	for i, v := range data {
		fmt.Println(i, v)
	}

	// while loop

	c := 0

	for {
		if c > 10 {
			break
		}

		// Your Code Here
		fmt.Println("while loop", c)

		// End Code
		c++
	}
}
