package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main1() {
	x := 10
	y := 20

	fmt.Println("До обмена:")
	fmt.Printf("x = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Println("После обмена:")
	fmt.Printf("x = %d, y = %d\n", x, y)
}
