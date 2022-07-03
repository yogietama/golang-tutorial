package main

import "fmt"

func main() {
	var outerScope int = 10
	if outerScope == 10 {
		innerScope := 20
		outerScope := 200 // inner scoped var
		outerScope = 100

		fmt.Println("inside if innerScope:", innerScope)
		fmt.Println("inside if outerScope:", outerScope)
	}

	for {
		innerScope2 := 50
		outerScope = 500
		fmt.Println("inside for loop innserScope2:", innerScope2)
		fmt.Println("inside for loop outerScope:", outerScope)
		break
	}

	fmt.Println("outside:", outerScope)
}
