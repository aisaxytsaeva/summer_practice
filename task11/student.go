package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Course int
	GPA    float64
}

func NewStudent(name string, age, course int, gpa float64) Student {
	return Student{
		Name:   name,
		Age:    age,
		Course: course,
		GPA:    gpa,
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.Age)
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.GPA)
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func mai1n() {
	student := NewStudent("Иван Петров", 20, 2, 4.5)

	student.PrintInfo()

	student.Promote()

	student.GPA = 4.7
	fmt.Printf("Новый средний балл: %.2f\n", student.GPA)
}
