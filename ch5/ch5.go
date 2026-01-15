package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Chapter 5")

	res, rem, _error := multipleReturnValuesFunc(5, 2)
	if _error != nil {
		fmt.Printf("%s\n", _error.Error())
	} else {
		fmt.Println(res, rem)
	}

	mult2 := makeMultFunc(2)

	fmt.Println(mult2(6)) // print 2 * 6
}

// Go supports multiple return values sort of like python (yay)
func multipleReturnValuesFunc(num int, denom int) (_ int, _ int, errorMsg error) {
	if denom == 0 {
		// Can predeclare return value (errorMsg)
		errorMsg = errors.New("cannot divide by 0")
		return 0, 0, errorMsg
	}
	return num / denom, num % denom, nil // return division res, remainder, and nil for error field
}

// Worth noting: functions are variables and can be treates as such, and
// function types can be defined, e.g. type FuncType func(int,int) int.
// Additionally, can use anonymous function within functions, called closures.

// Functions can return closures (inner functions)
func makeMultFunc(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
