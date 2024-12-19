package structs

import "fmt"

type CarEngine interface {
	TypeEngine() string
}

type Race struct{}

func (r Race) TypeEngine() string {
	return "V8"
}

type Popular struct{}

func (p Popular) TypeEngine() string {
	return "1.0"
}

func DescripCar(c CarEngine) {
	switch c.(type) {
	case Race:
		fmt.Println("This is a Race Car!")
	case Popular:
		fmt.Println("this is a popular Car!")
	default:
		fmt.Println("Unknow type")
	}
}

func SwitchStatement() {
	var car1 CarEngine = Race{}
	var car2 CarEngine = Popular{}

	DescripCar(car1)
	DescripCar(car2)

}
