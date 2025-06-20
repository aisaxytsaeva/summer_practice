package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int // Изменим на год рождения для вычисления возраста
	Course    int
	GPA       float64
}

func NewStudent(name string, birthYear, course int, gpa float64) Student {
	return Student{
		Name:      name,
		BirthYear: birthYear,
		Course:    course,
		GPA:       gpa,
	}
}

func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	switch {
	case s.GPA >= 4.5:
		return "отличник"
	case s.GPA >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.GetAge())
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.GPA)
	fmt.Printf("Статус: %s\n", s.GetStatus())
}

func main1() {
	student := NewStudent("Иван Петров", 2000, 2, 4.7)
	student.PrintInfo()
}
