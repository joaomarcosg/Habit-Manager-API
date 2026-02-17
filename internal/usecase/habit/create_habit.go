package habit

import (
	"context"
	"time"

	"github.com/joaomarcosg/Habit-Manager-API/internal/domain"
	"github.com/joaomarcosg/Habit-Manager-API/internal/validator"
)

type CreateHabitReq struct {
	Name        string           `json:"habit_name"`
	Category    string           `json:"habit_category"`
	Description string           `json:"description"`
	Frequency   []domain.WeekDay `json:"frequency"`
	StartDate   time.Time        `json:"start_date"`
	TargetDate  time.Time        `json:"target_date"`
	Priority    int              `json:"priority"`
}

func (req CreateHabitReq) Valid(ctx context.Context) validator.Evaluator {
	var eval validator.Evaluator

	eval.CheckField(validator.NotBlank(req.Name), "habit_name", "this field cannot be empty")
	eval.CheckField(validator.MaxChars(req.Name, 50), "habit_name", "habit name must be less than 50 chars")
	eval.CheckField(validator.NotBlank(req.Category), "habit_category", "this field cannot be empty")
	eval.CheckField(validator.NotBlank(req.Description), "habit_category", "this field cannot be empty")
	eval.CheckField(validator.MaxChars(req.Description, 150), "description", "this field must be less than 150 chars")
	eval.CheckField(validator.MaxLevel(req.Priority), "priority", "this field must be less than 10")
	eval.CheckField(validator.MinLevel(req.Priority), "priority", "this field must be greater than 0")

	return eval
}
