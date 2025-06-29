package main

import (
	"fmt"
)

type Engine struct {
	Type    string  // тип двигателя (бензиновый, дизельный, электрический)
	Power   float64 // мощность в л.с.
	Volume  float64 // объем в литрах
	IsTurbo bool    // наличие турбо
}

// Метод для вывода информации о двигателе
func (e Engine) PrintEngineInfo() {
	fmt.Printf("Тип двигателя: %s\n", e.Type)
	fmt.Printf("Мощность: %.1f л.с.\n", e.Power)
	fmt.Printf("Объем: %.1f л\n", e.Volume)
	fmt.Printf("Турбо: %t\n", e.IsTurbo)
}

// Структура Автомобиль
type Car struct {
	Brand      string
	Model      string
	Year       int
	Mileage    int    // пробег в км
	EngineInfo Engine // вложенная структура Двигатель
}

// Метод для вывода информации об автомобиле
func (c Car) PrintCarInfo() {
	fmt.Println("\n=== Информация об автомобиле ===")
	fmt.Printf("Марка: %s\n", c.Brand)
	fmt.Printf("Модель: %s\n", c.Model)
	fmt.Printf("Год выпуска: %d\n", c.Year)
	fmt.Printf("Пробег: %d км\n", c.Mileage)
	fmt.Println("\n--- Двигатель ---")
	c.EngineInfo.PrintEngineInfo()
	fmt.Println("==================")
}

// Метод для обновления пробега
func (c *Car) UpdateMileage(newMileage int) {
	if newMileage >= c.Mileage {
		c.Mileage = newMileage
	} else {
		fmt.Println("Ошибка: новый пробег не может быть меньше текущего")
	}
}

// Метод для проверки, является ли автомобиль спортивным
func (c Car) IsSportCar() bool {
	return c.EngineInfo.Power > 300 && c.EngineInfo.IsTurbo
}

func main() {
	engine := Engine{
		Type:    "бензиновый",
		Power:   450,
		Volume:  4.4,
		IsTurbo: true,
	}

	car := Car{
		Brand:      "BMW",
		Model:      "M5 Competition",
		Year:       2022,
		Mileage:    15000,
		EngineInfo: engine,
	}

	car.PrintCarInfo()

	if car.IsSportCar() {
		fmt.Println("\nЭто спортивный автомобиль!")
	} else {
		fmt.Println("\nЭто обычный автомобиль")
	}

	car.UpdateMileage(18000)
	fmt.Printf("\nОбновленный пробег: %d км\n", car.Mileage)

	car.UpdateMileage(17000)

	car.PrintCarInfo()
}
