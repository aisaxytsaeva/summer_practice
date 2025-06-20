package main

import "fmt"

type Engine struct {
	Power    int
	Volume   float64
	FuelType string
}

type Car struct {
	Make   string
	Model  string
	Year   int
	Engine Engine
}

func NewCar(make, model string, year int, power int, volume float64, fuelType string) Car {
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
		Engine: Engine{
			Power:    power,
			Volume:   volume,
			FuelType: fuelType,
		},
	}
}

func (c Car) PrintInfo() {
	fmt.Printf("Автомобиль: %s %s (%d год)\n", c.Make, c.Model, c.Year)
	fmt.Printf("Двигатель: %.1f л, %d л.с., топливо: %s\n",
		c.Engine.Volume, c.Engine.Power, c.Engine.FuelType)
}

func main() {
	car := NewCar("Toyota", "Camry", 2022, 248, 2.5, "бензин")

	car.PrintInfo()

	car.Engine.Power = 268
	fmt.Printf("Обновленная мощность: %d л.с.\n", car.Engine.Power)
}
