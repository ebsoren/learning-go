package main

import "fmt"

/*
	Many ways to create a Go array

	var x[3]int : len 3 int arr

	var x = [3]int{10,20,30}

	Go arrays support comparison w/ == and !=

	2D arr: var x [2][3]int

	Go considers arrays to consider length when typing, i.e. [2]int is not the same type as [3]int.
	Arrays are not used explicitly often for this reason. Also, cannot declare an array's length with a variable,
	since type needs to be known at compile time.

	Other consequences: can't write functions which take variable length arrays, can't assign variable to differently sized arrays.
	Arrays mainly exist to support slices, which essentially? function as arrays without the above typing restriction.
	However, comparison with == and != is not supported.

	i.e. var x = []int{1,2,3}
	[]int is a slice, [...]int is an array

*/

func main() {

	// var x = [5]int{10, 2: 20, 4: 30}

	y := []int{1, 2}
	x := []int{4, 5, 6}

	y = append(y, 3)

	// append x to y with ... operator
	y = append(y, x...)

	fmt.Println(len(y))
	fmt.Println(cap(y))

	for _, elt := range y {
		fmt.Printf("%d ", elt)
	}
	fmt.Printf("\n")

	z := make([]int, 6, 12)

	for _, elt := range z {
		fmt.Printf("%d ", elt)
	}
	fmt.Printf("\n")

	// Taking a slice from a slice means they share memory (so probably not relevant for
	// Distributed Systems)

	var x2 = x[:1]

	// 4 4 10 10
	fmt.Println(x2[0])
	fmt.Println(x[0])
	x2[0] = 10
	fmt.Println(x2[0])
	fmt.Println(x[0])

	// Compiler cannot detect if an array is created from a slice with a length larger than
	// the size of the slice - will panic at runtime.

	// Maps are written as var mapName[keyType]valueType

	// var totalWins map[string]int <-- Nilmap

	totalWins := map[string]int{} // <-- 0-map

	v, ok := totalWins["Esben"]

	fmt.Println(v, ok) // ok == false as key does not exist.

}
