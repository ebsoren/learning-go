package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Chapter 6")

	// One use for pointers: to/from JSON

	f := struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}{}
	err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Name: %s, Age: %d\n", f.Name, f.Age)

}
