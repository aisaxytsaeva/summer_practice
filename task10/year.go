package main

import "fmt"

func isLeapYear(year int) bool {
	return year%400 == 0 || (year%100 != 0 && year%4 == 0)
}

func main() {
	var year int

	fmt.Print("Enter year: ")
	if _, err := fmt.Scanln(&year); err != nil {
		fmt.Println("Invalid year input")
		return
	}

	if isLeapYear(year) {
		fmt.Printf("%d is a leap year\n", year)
	} else {
		fmt.Printf("%d is NOT a leap year\n", year)
	}
}
