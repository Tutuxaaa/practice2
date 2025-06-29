package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int // Изменим Age на BirthYear для вычисления возраста
	Course    int
	AvgGrade  float64
}

// Функция создания нового студента
func NewStudent(name string, birthYear, course int, avgGrade float64) Student {
	return Student{
		Name:      name,
		BirthYear: birthYear,
		Course:    course,
		AvgGrade:  avgGrade,
	}
}

// Функция вывода информации о студенте
func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\n", s.Name)
	fmt.Printf("Возраст: %d\n", s.GetAge())
	fmt.Printf("Курс: %d\n", s.Course)
	fmt.Printf("Средний балл: %.2f\n", s.AvgGrade)
	fmt.Printf("Статус: %s\n", s.GetStatus())
}

// Функция проверки стипендии
func (s Student) HasScholarship() bool {
	return s.AvgGrade >= 4.5
}

// Функция изменения среднего балла
func (s *Student) UpdateGrade(newGrade float64) {
	if newGrade >= 0 && newGrade <= 5 {
		s.AvgGrade = newGrade
	} else {
		fmt.Println("Ошибка: средний балл должен быть между 0 и 5")
	}
}

// Метод для вычисления возраста (по году рождения)
func (s Student) GetAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

// Метод для получения статуса студента
func (s Student) GetStatus() string {
	switch {
	case s.AvgGrade >= 4.8:
		return "отличник"
	case s.AvgGrade >= 3.8:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student := NewStudent("Иван Иванов", 2000, 2, 4.7)

	student.PrintInfo()

	if student.HasScholarship() {
		fmt.Println("Студент получает стипендию")
	} else {
		fmt.Println("Студент не получает стипендию")
	}
	fmt.Println()

	student.UpdateGrade(4.9)
	student.PrintInfo()

	student.UpdateGrade(5.1)
	fmt.Println()

	// Демонстрация работы новых методов
	fmt.Printf("Текущий возраст: %d\n", student.GetAge())
	fmt.Printf("Текущий статус: %s\n", student.GetStatus())
}
