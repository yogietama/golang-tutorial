// goplay.space/#dy-g6N_H3VJ
package main

import (
	"fmt"
)

func main() {
	var emptyInterface interface{}
	emptyInterface = "alpha"
	emptyInterface = 25

	// trying to get string value
	// cast to string
	stringVal, ok := emptyInterface.(string)
	if !ok {
		fmt.Println("emptyInterface does not contain string value")
	}
	fmt.Println(stringVal == "") // default val

	// trying to get int value
	// cast to int
	intVal, ok := emptyInterface.(int)
	if !ok {
		fmt.Println("emptyInterface does not contain int value")
	}
	fmt.Println(intVal) // 25
	fmt.Println(emptyInterface)
}
