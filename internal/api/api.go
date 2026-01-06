package api

import "github.com/joaomarcosg/Habit-Manager-API/internal/services"

type Api struct {
	UserService services.UserService
}
