package main

import "fmt"

func main1() {
	var num1, num2 float64

	fmt.Print("Enter first number: ")
	if _, err := fmt.Scanln(&num1); err != nil {
		fmt.Println("Invalid input for first number")
		return
	}

	fmt.Print("Enter second number: ")
	if _, err := fmt.Scanln(&num2); err != nil {
		fmt.Println("Invalid input for second number")
		return
	}

	fmt.Printf("\n%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
	fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
	fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, num1*num2)

	if num2 != 0 {
		fmt.Printf("%.2f / %.2f = %.2f\n", num1, num2, num1/num2)
	} else {
		fmt.Println("Cannot divide by zero")
	}

	int1, int2 := int(num1), int(num2)
	if int2 != 0 {
		fmt.Printf("%d %% %d = %d (remainder)\n", int1, int2, int1%int2)
	} else {
		fmt.Println("Cannot calculate remainder with zero")
	}
}
