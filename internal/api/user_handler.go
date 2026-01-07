package api

import (
	"errors"
	"net/http"

	"github.com/joaomarcosg/Habit-Manager-API/internal/jsonutils"
	"github.com/joaomarcosg/Habit-Manager-API/internal/services"
	"github.com/joaomarcosg/Habit-Manager-API/internal/usecase/user"
)

func (api *Api) handleSignupUser(w http.ResponseWriter, r *http.Request) {

	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.UserService.CreateUser(
		r.Context(),
		data.Name,
		data.Email,
		data.Password,
	)

	if err != nil {
		if errors.Is(err, services.ErrDuplicatedEmailOrUserName) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or username already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"user_id": id,
	})

}
