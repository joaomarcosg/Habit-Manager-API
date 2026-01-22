package api

import (
	"errors"
	"net/http"

	"github.com/joaomarcosg/Habit-Manager-API/internal/jsonutils"
	"github.com/joaomarcosg/Habit-Manager-API/internal/services"
	"github.com/joaomarcosg/Habit-Manager-API/internal/usecase/category"
)

func (api *Api) handleCreateCategory(w http.ResponseWriter, r *http.Request) {

	data, problems, err := jsonutils.DecodeValidJson[category.CreateCategoryReq](r)

	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
		return
	}

	id, err := api.CategoryService.CreateCategory(r.Context(), data.Name)

	if err != nil {
		if errors.Is(err, services.ErrDuplicateCategoryName) {
			_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, map[string]any{
				"error": "email or username already exists",
			})
			return
		}
	}

	_ = jsonutils.EncodeJson(w, r, http.StatusCreated, map[string]any{
		"category_id": id,
	})

}
