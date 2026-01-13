package main

import (
	"fmt"
	"math/rand"
)

/*
Go uses variable shadowing, can even shadow package names, built in types, and constants.
For example, can shadow true, false, make, close, nil, int, string...
*/
func main() {
	fmt.Println("Chapter 4")

	// n scoped in the if statement chain
	if n := rand.Intn(10); n == 0 {
		fmt.Println("n is so low")
	} else if n > 5 {
		fmt.Println("n is so high")
	} else {
		fmt.Println("n is perfect")
	}

	/*
		4 ways to do for loop: C-style, condition-only, infinite, for-range
	*/

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\n")

	i := 0

	for i < 10 {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Printf("\n")

	// Infinite
	/*
			func main() {
		    	for {
		    	    fmt.Println("Hello")
		    	}
			}
	*/

	evenVals := []int{2, 4, 6, 8}

	// NOTE: for-range creates copies, cannot modify in loop
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// switch statement
	words := []string{"one", "three", "five", "twelve", "one hundred", "three"}

outer:
	for _, v := range words {
		switch len(v) {
		case 0, 1, 2, 3:
			fmt.Println(v, "is short")
		case 4, 5, 6:
			fmt.Println(v, "is middle")

		default:
			fmt.Println(v, "is long. breaking")
			break outer

		}
	}

	intSlice := []int{}
	for i := 0; i <= 100; i++ {
		intSlice = append(intSlice, i)
	}

	for _, v := range intSlice {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")

}
