package main

import "fmt"

func main1() {
	var number1, number2 float64
	var operation string

	fmt.Print("Enter number1: ")
	fmt.Scanln(&number1)

	fmt.Print("Enter number2: ")
	fmt.Scanln(&number2)

	fmt.Print("Enter operation: ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Printf("Result: %.2f\n", number1+number2)
	case "-":
		fmt.Printf("Result: %.2f\n", number1-number2)
	case "*":
		fmt.Printf("Result: %.2f\n", number1*number2)
	case "/":
		if number2 == 0 {
			fmt.Println("Error: division by zero")
		} else {
			fmt.Printf("Result: %.2f\n", number1/number2)
		}
	}

}
