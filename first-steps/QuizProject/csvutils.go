package quizproject

import (
	"errors"
	"strconv"
)

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("letters are not allowed")
	}
	return i, nil
}
