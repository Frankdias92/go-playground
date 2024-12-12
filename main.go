package main

import (
	"fmt"
)

type Car struct {
	Name string
}

func (c Car) move() {
	c.Name = "Audi"
}

func (ptr *Car) edit() {
	ptr.Name = "Mercedez"
}

func main() {
	car := Car{
		Name: "Mustang",
	}
	car.move()
	fmt.Println(car.Name)

	car.edit()
	fmt.Println(car.Name)

	// points.GetPoints()
	// points.Exercice()

}
