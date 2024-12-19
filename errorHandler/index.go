package errorhandler

import (
	"errors"
	"fmt"
)

func HandleError() {
	a := 10
	b := 0

	res, err := Divid(a, b)
	if err != nil {
		fmt.Printf("res: %v\n", err)
		return
	}
	fmt.Println(a, "/", b, res)
}

func Divid(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("you cant divid by 0")
	}
	return a / b, nil
}
