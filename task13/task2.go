package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func incrementAgeByValue(p Person) {
	p.Age++
	fmt.Println("Внутри incrementAgeByValue:", p.Age)
}

func incrementAgeByPointer(p *Person) {
	p.Age++
	fmt.Println("Внутри incrementAgeByPointer:", p.Age)
}

func main() {
	person := Person{"Иван", 30}

	fmt.Println("Оригинал перед вызовами:", person)

	incrementAgeByValue(person)
	fmt.Println("После incrementAgeByValue:", person)

	incrementAgeByPointer(&person)
	fmt.Println("После incrementAgeByPointer:", person)

	num := 5
	passByValue(num)
	fmt.Println("num после passByValue:", num) // 5

	passByPointer(&num)
	fmt.Println("num после passByPointer:", num) // 10
}

func passByValue(n int) {
	n = 10
}

func passByPointer(n *int) {
	*n = 10
}
