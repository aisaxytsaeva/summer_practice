package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Task1")
	grades := make(map[string]int)
	addGrade(grades, "Ivan", 5)
	addGrade(grades, "Till", 4)
	addGrade(grades, "Luka", 3)

	fmt.Println("Ivan's garades:", getGrade(grades, "Ivan"))

	removeGrade(grades, "Ivan")
	fmt.Println("After deleting:", grades)

	fmt.Println("Task2")
	str := "thank you for being the victim of my shallow emotions thank you for being the victim of my shallow emotions thank you for being the victim the victim ivan till mizi"
	freq := wordFrequency(str)

	fmt.Println("Word freq:")
	for word, count := range freq {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func wordFrequency(text string) map[string]int {
	words := strings.Fields(strings.ToLower(text))
	freq := make(map[string]int)

	for _, word := range words {
		freq[word]++
	}

	return freq
}

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
}

func getGrade(grades map[string]int, name string) int {
	return grades[name]
}

func removeGrade(grades map[string]int, name string) {
	delete(grades, name)
}
