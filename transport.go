package main

import "fmt"

type Transport interface {
	Move()        // Движение
	Stop()        // Остановка
	Name() string // Название транспорта
}

//структура автомобиля
type Car struct {
	model string
	speed int
}

func (c Car) Move() {
	fmt.Printf("%s едет со скоростью %d км/ч\n", c.model, c.speed)
}

func (c Car) Stop() {
	fmt.Printf("%s остановился\n", c.model)
}

func (c Car) Name() string {
	return c.model
}

//структура велосипеда
type Bicycle struct {
	brand string
}

func (b Bicycle) Move() {
	fmt.Printf("Велосипед %s крутит педали\n", b.brand)
}

func (b Bicycle) Stop() {
	fmt.Printf("Велосипед %s остановился\n", b.brand)
}

func (b Bicycle) Name() string {
	return "Велосипед " + b.brand
}

// Airplane - структура самолёта
type Airplane struct {
	flightNumber string
	altitude     int
}

func (a Airplane) Move() {
	fmt.Printf("Самолёт рейса %s набирает высоту %d метров\n", a.flightNumber, a.altitude)
}

func (a Airplane) Stop() {
	fmt.Printf("Самолёт рейса %s приземлился\n", a.flightNumber)
}

func (a Airplane) Name() string {
	return "Самолёт " + a.flightNumber
}

// TransportControl - функция для управления любым транспортом
func TransportControl(t Transport) {
	fmt.Printf("\nУправляем транспортом: %s\n", t.Name())
	t.Move()
	t.Stop()
}

func main() {
	// Создаем разные виды транспорта
	car := Car{model: "Toyota Camry", speed: 60}
	bike := Bicycle{brand: "Stels"}
	plane := Airplane{flightNumber: "SU-137", altitude: 10000}

	// Управляем транспортом через интерфейс
	TransportControl(car)
	TransportControl(bike)
	TransportControl(plane)

	// Можно также работать с коллекцией транспорта
	transports := []Transport{car, bike, plane}

	fmt.Println("\nДемонстрация работы всего транспорта:")
	for _, t := range transports {
		t.Move()
		t.Stop()
		fmt.Println()
	}
}
