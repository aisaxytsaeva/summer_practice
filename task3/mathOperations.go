package main

import (
	"fmt"
	"math"
)

func main() {

	const pi = math.Pi
	const e = math.E

	circleRadius := 5.0
	circleArea := pi * math.Pow(circleRadius, 2)
	exponentialResult := math.Pow(e, 2)

	fmt.Printf("Значение π (пи): %.4f\n", pi)
	fmt.Printf("Значение e (экспонента): %.4f\n", e)
	fmt.Printf("Площадь круга с радиусом %.1f: %.2f\n", circleRadius, circleArea)
	fmt.Printf("e^2 ≈ %.4f\n", exponentialResult)
}
