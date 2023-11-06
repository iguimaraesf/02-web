package main

import (
	"fmt"

	"github.com/ivanguimaraes/goweb/packages/car"
)

func main() {
	car := car.Car{"Gol", "Black"}
	fmt.Println(car.Start())
}
