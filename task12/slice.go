package main

import (
	"fmt"
	"sort"
)

// serch for element
func FindElement(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

// sort
func SortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

// filter
func FilterEvenNumbers(slice []int) []int {
	var result []int
	for _, v := range slice {
		if v%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main1() {
	numbers := []int{5, 2, 8, 1, 9, 3, 6, 4}

	if index, found := FindElement(numbers, 8); found {
		fmt.Printf("Элемент 8 найден на позиции %d\n", index)
	} else {
		fmt.Println("Элемент не найден")
	}

	sorted := SortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	evenNumbers := FilterEvenNumbers(numbers)
	fmt.Println("Четные числа:", evenNumbers)
}
