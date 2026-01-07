package user

import (
	"context"

	"github.com/joaomarcosg/Habit-Manager-API/internal/validator"
)

type CreateUserReq struct {
	Name     string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (req CreateUserReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Name), "user_name", "this field cannot be empty")
	eval.CheckField(validator.NotBlank(req.Email), "email", "this field cannot be empty")
	eval.CheckField(validator.Matches(req.Email, validator.EmailRX), "email", "must be a valid email")
	eval.CheckField(validator.MinChars(req.Password, 8), "password", "password must be bigger than 8 chars")

	return eval
}
