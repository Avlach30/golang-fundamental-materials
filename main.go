package main

import "fmt"

func main() {
	var name string //* Declarating empty variable with string data type

	name = "Ahmad Muzakki" //* Reassign existing variable
	fmt.Println(name)
	name = "Ahmad Sujancok"
	fmt.Println(name)
	println("========")
	var address = "Diponegoro street 5" //* Declarating variable without data type
	fmt.Println(address)
	address = "Pattimura 77"
	fmt.Println(address)
	println("========")
	country := "Indonesia" //* Declarating varriable withour var keyword and data type
	fmt.Println(country)
	country = "Iceland"
	println(country)
	println("=========")
	//* Declarating multiple variable directly
	var (
		student1 = "Andra"
		student2 = "Bobby"
	)
	fmt.Println(student1)
	fmt.Println(student2)
	println("*********")
	//* Declarating constant in go, constant value can't reassign
	const phi = 3.14
	fmt.Println(phi)
	println("========")
	//* Declarating multiple constant directly
	const (
		circleRadius float32 = 12
		circleWide = phi * (circleRadius * circleRadius)
	)
	fmt.Println(circleRadius)
	fmt.Println(circleWide)
}