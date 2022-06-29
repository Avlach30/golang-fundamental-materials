package main

import "fmt"

func main() {
	var name string //* Declarating empty variable with string data type

	name = "Ahmad Muzakki" //* Reassign existing variable
	fmt.Println(name)
	println("========")
	name = "Ahmad Sujancok"
	fmt.Println(name)

	var address = "Diponegoro street 5" //* Declarating variable without data type
	fmt.Println(address)
	address = "Pattimura 77"
	fmt.Println(address)
}