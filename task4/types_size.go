package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Целочисленные типы
	var intVar int = 42
	var int8Var int8 = 127
	var int16Var int16 = 32767
	var int32Var int32 = 2147483647
	var int64Var int64 = 9223372036854775807

	// Числа с плавающей точкой
	var float32Var float32 = 3.1415927
	var float64Var float64 = 3.141592653589793

	// Комплексные числа
	var complex64Var complex64 = 1 + 2i
	var complex128Var complex128 = 1 + 2i

	// Логический тип
	var boolVar bool = true

	// Строки
	var stringVar string = "Hello, User!"

	fmt.Println("Целочисленные:")
	fmt.Printf("int: %d, int8: %d, int16: %d, int32: %d, int64: %d\n",
		intVar, int8Var, int16Var, int32Var, int64Var)

	fmt.Println("\nДробные:")
	fmt.Printf("float32: %g, float64: %g\n", float32Var, float64Var)

	fmt.Println("\nКомплексные:")
	fmt.Printf("complex64: %v, complex128: %v\n", complex64Var, complex128Var)

	fmt.Println("\nПрочие:")
	fmt.Printf("bool: %t, string: %s\n", boolVar, stringVar)

	fmt.Println("Размеры типов данных (в байтах):")

	// Целочисленные
	fmt.Printf("int: %d\n", unsafe.Sizeof(intVar))
	fmt.Printf("int8: %d\n", unsafe.Sizeof(int8Var))
	fmt.Printf("int16: %d\n", unsafe.Sizeof(int16Var))
	fmt.Printf("int32: %d\n", unsafe.Sizeof(int32Var))
	fmt.Printf("int64: %d\n", unsafe.Sizeof(int64Var))

	// Дробные
	fmt.Printf("float32: %d\n", unsafe.Sizeof(float32Var))
	fmt.Printf("float64: %d\n", unsafe.Sizeof(float64Var))

	// Комплексные
	fmt.Printf("complex64: %d\n", unsafe.Sizeof(complex64Var))
	fmt.Printf("complex128: %d\n", unsafe.Sizeof(complex128Var))

	// Логические
	fmt.Printf("bool: %d\n", unsafe.Sizeof(boolVar))

	// Строки
	fmt.Printf("string: %d (структура)\n", unsafe.Sizeof(stringVar))

}
