package errorhandler

import (
	"errors"
	"fmt"
)

func Error1() {

	user, err := NewUser(true)
	if err != nil {
		fmt.Println("something wrong on create user")
		return
	}
	println(user)
	// user.Foo() /* this wil result a panic */
}

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("an error")
	}
	return &User{}, nil
}
