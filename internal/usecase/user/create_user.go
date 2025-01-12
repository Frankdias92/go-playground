package user

import (
	"context"

	"github.com/frankdias92/go-playground/internal/usecase/validator"
)

type CreateUser struct {
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"password_hash"`
	Bio          string `json:"bio"`
}

func (req CreateUser) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NoBlank(req.UserName), "user_name", "this field cannot be empty")

	return eval
}
