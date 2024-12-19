package errorhandler

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func Error3() {
	err := foo()
	var sqrtError *SqrtError

	if err != nil && errors.As(err, &sqrtError) {
		fmt.Println(sqrtError.msg)
		return
	}
	fmt.Println("out side message")
}

func foo() error { return &SqrtError{msg: "test"} }
