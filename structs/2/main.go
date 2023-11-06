package main

import "fmt"

type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	Name   string
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corolla", 2017, "White"}
	sCar := SuperCar{
		Car: Car{
			"Fusca", 2037, "Blue",
		},
		CanFly: true,
		Name:   "Elantra"}
	fmt.Println(car1)
	fmt.Println(sCar.Name, sCar.Car.Name)
	fmt.Println(sCar.info())
}
