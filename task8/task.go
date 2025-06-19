package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter number from 1 to 10: ")
	fmt.Scanln(&number)

	if number < 1 || number > 10 {
		fmt.Println("Error: Enter number from 1 to 10")
		return
	}

	fmt.Printf("\nmultiplication table for %d:\n", number)
	fmt.Println("-------------------")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
	fmt.Println("-------------------")

	fmt.Println("Simple numbers till 100:")

	for num := 2; num <= 100; num++ {
		isPrime := true

		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			fmt.Printf("%d ", num)
		}
	}
}
