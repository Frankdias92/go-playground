package errorhandler

import (
	"errors"
	"fmt"
)

func Error5() {
	err := foo5()
	fmt.Println(err)
	fmt.Println(errors.Is(err, ErrSomething2))
}

var ErrSomething2 = errors.New("error")

func a() error { return ErrSomething2 }
func b() error { return ErrSomething2 }

func foo5() error {
	var errorResult error

	if err := a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	return errorResult
}
