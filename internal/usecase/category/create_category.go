package category

import (
	"context"

	"github.com/joaomarcosg/Habit-Manager-API/internal/validator"
)

type CreateCategoryReq struct {
	Name string `json:"category_name"`
}

func (req CreateCategoryReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Name), "category_name", "this field cannot be empty")
	eval.CheckField(validator.MaxChars(req.Name, 50), "category_name", "category name must be less than 50 chars")

	return eval
}
