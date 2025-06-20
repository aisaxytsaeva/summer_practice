package main

import "fmt"

type Transport interface {
	Move(speed int)
	Stop()
}

type Car struct {
	Model string
}

func (c Car) Move(speed int) {
	fmt.Printf("Автомобиль %s едет со скоростью %d км/ч\n", c.Model, speed)
}

func (c Car) Stop() {
	fmt.Printf("Автомобиль %s остановился\n", c.Model)
}

type Bicycle struct {
	Type string
}

func (b Bicycle) Move(speed int) {
	fmt.Printf("Велосипед (%s) движется со скоростью %d км/ч\n", b.Type, speed)
}

func (b Bicycle) Stop() {
	fmt.Printf("Велосипед (%s) остановился\n", b.Type)
}

func testTransport(t Transport) {
	t.Move(60)
	t.Stop()
	fmt.Println("---")
}

func main() {
	car := Car{Model: "Toyota Camry"}
	bike := Bicycle{Type: "Горный"}

	fmt.Println("Тестируем транспорт:")
	testTransport(car)
	testTransport(bike)
}
