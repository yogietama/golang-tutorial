package main

import "fmt"

// func with separate type declaration
func Concat(x string, y string) string {
	return x + y
}

// func with group same type declaration
func Concat2(x, y string) string {
	return x + y
}

// func with multiple return
func Concat3(x, y string) (string, error) {
	if x == "" || y == "" {
		return "", fmt.Errorf("Params cannot be empty")
	}

	return x + y, nil
}

// variadic functions
func add(x ...int) int {
	temp := 0
	for _, v := range x {
		temp += v
	}

	return temp
}

func main() {
	res := Concat("aplha", "beta")
	res2 := Concat2("alpha", "beta")
	res3, err := Concat3("alpha", "")
	total := add(1, 3, 4, 5)
	fmt.Println("Concat result  :", res)
	fmt.Println("Concat2 result :", res2)
	fmt.Println("Concat3 result :", res3, err)
	fmt.Println("Add Result :", total)
}
