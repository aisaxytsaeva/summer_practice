package main

import "fmt"

func main() {
	var dayNumber int

	fmt.Print("Enter a number (1-7): ")
	fmt.Scanln(&dayNumber)

	switch dayNumber {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Error")
	}
}
