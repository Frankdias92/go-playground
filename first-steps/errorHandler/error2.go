package errorhandler

import (
	"fmt"
	"math"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

func sqroot(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"there is no such sqroot of negative value"}
	}
	result := math.Sqrt(x)
	return result, nil
}

func Error2() {
	x := -10
	res, err := sqroot(float64(x))

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
