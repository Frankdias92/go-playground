package errorhandler

import (
	"errors"
	"fmt"
)

func Error4() {
	err := foo4()

	if err != nil && errors.Is(err, ErrSomething) {
		fmt.Println("get an error: ", err)
		return
	}
}

func foo4() error {
	err := bar()

	if err != nil {
		return fmt.Errorf("something wrong in foo: %w", err)
	}
	return nil
}

var ErrSomething = errors.New("error")

func bar() error {
	return ErrSomething
}
