package api

import (
	"net/http"

	"github.com/joaomarcosg/Habit-Manager-API/internal/jsonutils"
	"github.com/joaomarcosg/Habit-Manager-API/internal/usecase/habit"
)

func (api *Api) handleCreateHabit(w http.ResponseWriter, r *http.Request) {

	data, problems, err := jsonutils.DecodeValidJson[habit.CreateHabitReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.HabitService.CreateHabit(
		r.Context(),
		data.Name,
		data.Category,
		data.Description,
		data.Frequency,
		data.StartDate,
		data.TargetDate,
		data.Priority,
	)

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"habit_id": id,
	})

}
