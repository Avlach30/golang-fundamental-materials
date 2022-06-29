package main

import "fmt"

func main() {
	firstNum := 10
	secondNum := 18

	//* Define math operations
	fmt.Println(firstNum + secondNum)
	fmt.Println(secondNum % firstNum)
	fmt.Println(-10) //* Define negative number value
	println("======")
	//* Define comparation operator, returned bool value
	fmt.Println(firstNum != secondNum)
	fmt.Println(firstNum == secondNum)
	println("=======")
	finalScore := 89
	absentScore := 86

	passFinalScore := finalScore >= 88
	passAbsentScore := absentScore >= 87

	//* Define logical (boolean operator)
	passResult := passFinalScore && passAbsentScore
	fmt.Println(passResult)
	advanceClassResult := passFinalScore || finalScore >= 80
	fmt.Println(advanceClassResult)
}