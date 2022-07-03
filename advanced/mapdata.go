package main

import (
	"fmt"
)

func main() {
	// init a map
	var myDictionary map[string]string
	nameAgeMap := make(map[string]int)
	nestedMap := make(map[int]map[string]string)

	fmt.Println(myDictionary, myDictionary == nil) // true
	fmt.Println(nameAgeMap, nameAgeMap != nil)     // true
	fmt.Println(nestedMap, nestedMap != nil)       // true

	// myDictionary["school"] = "place to study" --> error, panic: assignment to entry in nil map
	myDictionary = make(map[string]string)
	myDictionary["school"] = "place to study"
	nameAgeMap["nicky"] = 22
	nestedMap[12] = make(map[string]string)
	// without the above stmt, below stmt will panic
	nestedMap[12]["abc"] = "def"

	// how to check if key exists
	nickyAge, ok := nameAgeMap["nicky"]
	fmt.Println(nickyAge == 22) // true

	jamesAge, ok := nameAgeMap["james"]
	if !ok {
		fmt.Println("james' age not found")
	}
	fmt.Println(jamesAge == 0)

	fmt.Println(myDictionary)
	fmt.Println(nameAgeMap)
	fmt.Println(nestedMap)

	// access map with for range
	fmt.Println("===== Access Map with range =====")
	nameAgeRangeMap := make(map[string]int)
	nameAgeRangeMap["nicky"] = 18
	nameAgeRangeMap["jane"] = 28
	nameAgeRangeMap["alice"] = 13
	nameAgeRangeMap["jason"] = 43

	// nameAgeRange length = 4
	fmt.Println("nameAgeRange length:", len(nameAgeRangeMap))

	// No order Guarantee
	for k, v := range nameAgeRangeMap {
		fmt.Println(k, v)
	}

}
