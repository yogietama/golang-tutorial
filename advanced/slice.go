// goplay.space/#B74rJxMvEOR

package main

import (
	"fmt"
)

func main() {
	names := [4]string{"james", "jack", "jill", "ricky"}
	var someNames []string
	fmt.Println(someNames == nil) // true, default value
	// someNames[0] = "abc" -> PANIC
	someNames = append(someNames, "abc") // WORKS
	fmt.Println(someNames)               // output: [abc]

	// internally using 'names' array
	someNames = names[1:3]
	fmt.Println(len(someNames)) // len: 2
	fmt.Println(cap(someNames)) // cap: 3
	fmt.Println(someNames)      // output: [jack jill]

	// internally auto created underlying array
	someNewNames := []string{"alpha", "beta"}
	fmt.Println(len(someNewNames)) // len: 2
	fmt.Println(cap(someNewNames)) // cap: 2
	fmt.Println(someNewNames)      // output: [alpha beta]
	someNewNames = append(someNewNames, "theta")
	fmt.Println(someNewNames) // output: [alpha beta theta]

	// WARNING 1
	someNames[1] = "theta" // this changes the underlying array
	fmt.Println(someNames) // output: [jack theta]
	fmt.Println(names)     // output: [james jack theta ricky]

	// WARNING 2
	someNames = append(someNames, "beta")
	// stmt above also changes the underlying array
	fmt.Println(someNames) // output: [jack theta beta]
	fmt.Println(names)     // output: [james jack theta beta]

	// CAVEAT
	fmt.Println(len(someNames), cap(someNames)) // 3, 3
	someNames = append(someNames, "gamma")
	// stmt above will create a new underlying array
	// so it will not change the underlying array
	fmt.Println(someNames) // output: [jack theta beta gamma]
	fmt.Println(names)     // output: [james jack theta beta]

	// For range slice

	names1 := []string{"jack", "alice", "james"}

	// ORDERED
	for idx, val := range names1 {
		fmt.Println(idx, val)
	}
}
