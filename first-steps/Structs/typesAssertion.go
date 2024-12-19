package structs

import "fmt"

type Car interface {
	TypeCar() string
}

type Mercedez struct{}

func (m *Mercedez) TypeCar() string {
	return "v12 engine"
}

func WhatEngine(c Car) {
	if mercedez, ok := c.(*Mercedez); ok {
		fmt.Println("It's a", mercedez.TypeCar())
	} else {
		fmt.Println("this is not a v12")
	}
}

func TypeStruct() {
	var car Car = &Mercedez{}
	WhatEngine(car)

}
