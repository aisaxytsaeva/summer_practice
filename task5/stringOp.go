package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "thank you for being the victim of my shallow emotions"
	fmt.Println("The length of the string", len(str))
	var substr = "you"
	var isContain = strings.Contains(str, substr)
	fmt.Printf("Is %s in the string? %t\n", substr, isContain)

	fmt.Println("Upper String:", strings.ToUpper(str))
	fmt.Println("Lower String:", strings.ToLower(str))

	fmt.Println("Task2")

	words := strings.Split(str, " ")
	for _, word := range words {
		fmt.Println(word)
	}

}
