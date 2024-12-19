package structs

import "fmt"

type Animal interface {
	Soud() string
}

type Dog struct{}

func (Dog) Soud() string {
	return "Au! Au!"
}

func whatDoesThisAnimalSay(a Animal) {
	fmt.Println(a.Soud())
}

func Interface() {
	dog := Dog{}
	whatDoesThisAnimalSay(dog)
}
