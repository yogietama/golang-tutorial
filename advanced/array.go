// goplay.space/Nw5Oy0E6zfP

package main

import (
	"fmt"
)

func main() {
	var fruits [5]string
	fmt.Println(fruits == [5]string{}) // default value
	fruits = [5]string{"apple", "orange"}
	fmt.Println(len(fruits)) // len: 5

	ages := [3]int{12, 15, 18}
	names := [...]string{"jack", "jill", "janice"}
	fmt.Println(len(ages))  // len: 3
	fmt.Println(len(names)) // len: 3

	names[1] = "james"
	fmt.Println(names) // output: [jack james janice]

	// ORDERED
	for idx, val := range names {
		fmt.Println(idx, val)
	}
}
