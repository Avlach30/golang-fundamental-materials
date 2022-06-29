package main

import "fmt"

func main() {
	type marriedStatus bool //* Define custom data type

	//* Define variable with custom data type
	var isMarried marriedStatus = true
	fmt.Println(isMarried)
}