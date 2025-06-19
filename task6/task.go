package main

import "fmt"

func main() {
	var manga = []string{"Kagurabschi", "Sakamoto Days", "One piece", "D.Gray-man"}
	manga = append(manga, "Wind Breaker")
	manga = append(manga, "HQ!")

	fmt.Println("slice content:")

	for i, title := range manga {
		fmt.Printf("Index %d: %s\n", i, title)
	}

	manga = removeElement(manga, 2)

	fmt.Println("slice content after deleting:")

	for i, title := range manga {
		fmt.Printf("Index %d: %s\n", i, title)
	}
}

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
