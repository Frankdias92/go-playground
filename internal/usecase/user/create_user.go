package user

import (
	"context"

	"github.com/frankdias92/go-playground/internal/usecase/validator"
)

type CreateUser struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
	Bio      string `json:"bio"`
}

func (req CreateUser) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NoBlank(req.UserName), "user_name", "this field cannot be empty")
	eval.CheckField(validator.NoBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "must be an valid email")
	eval.CheckField(validator.NoBlank(req.Bio), "bio", "this field cannot be empty")
	eval.CheckField(
		validator.MinChars(req.Bio, 10) &&
			validator.MaxChars(req.Bio, 255), "bio", "this field must have a lenght between 10 asn 255",
	)

	eval.CheckField(validator.MinChars(string(req.Password), 8), "password", "must be bigger than 8 chars")

	return eval
}
