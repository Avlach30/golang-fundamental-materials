package main

import "fmt"

func main() {
	//* Converting string to another string
	helloStr := "Hello"
	fmt.Println(helloStr)
	lStr := string(helloStr[2]) //* Get single character from existing variable of string
	fmt.Println(lStr)
	println("========")
}